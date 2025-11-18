package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

// -------------------- Events --------------------

type Event interface {
	EventType() string
	OccurredAt() time.Time
}

type BaseEvent struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BaseEvent) EventType() string   { return e.Type }
func (e BaseEvent) OccurredAt() time.Time { return e.Timestamp }

// AccountOpened
type AccountOpened struct {
	BaseEvent
	AccountID string  `json:"account_id"`
	Owner     string  `json:"owner"`
	Balance   float64 `json:"balance"`
}

// MoneyDeposited
type MoneyDeposited struct {
	BaseEvent
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

// MoneyWithdrawn
type MoneyWithdrawn struct {
	BaseEvent
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

// -------------------- Event Store --------------------

// Simple append-only in-memory event store (thread-safe)
type EventStore struct {
	mu     sync.RWMutex
	events map[string][]Event // streamID -> events
}

func NewEventStore() *EventStore {
	return &EventStore{
		events: make(map[string][]Event),
	}
}

func (s *EventStore) Append(streamID string, ev Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[streamID] = append(s.events[streamID], ev)
}

func (s *EventStore) Load(streamID string) []Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	cpy := make([]Event, len(s.events[streamID]))
	copy(cpy, s.events[streamID])
	return cpy
}

// LoadAll returns all events in chronological appended order (simple)
func (s *EventStore) LoadAll() []Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	all := []Event{}
	for _, evs := range s.events {
		all = append(all, evs...)
	}
	return all
}

// -------------------- Aggregate (Account) --------------------

type Account struct {
	ID      string
	Owner   string
	Balance float64
	// event history not stored here in full in ES pattern; we manage via event store
}

func (a *Account) Apply(e Event) error {
	switch ev := e.(type) {
	case AccountOpened:
		a.ID = ev.AccountID
		a.Owner = ev.Owner
		a.Balance = ev.Balance
	case MoneyDeposited:
		a.Balance += ev.Amount
	case MoneyWithdrawn:
		a.Balance -= ev.Amount
	default:
		return fmt.Errorf("unknown event type %T", e)
	}
	return nil
}

// Rehydrate aggregate from a list of events
func RehydrateAccount(events []Event) (*Account, error) {
	acc := &Account{}
	for _, e := range events {
		if err := acc.Apply(e); err != nil {
			return nil, err
		}
	}
	if acc.ID == "" && len(events) == 0 {
		return nil, errors.New("no events")
	}
	return acc, nil
}

// -------------------- Commands & Handlers --------------------

type Command interface{}

type OpenAccountCmd struct {
	AccountID string
	Owner     string
	Initial   float64
}

type DepositCmd struct {
	AccountID string
	Amount    float64
}

type WithdrawCmd struct {
	AccountID string
	Amount    float64
}

// CommandHandler: perform business logic, emit events
type CommandHandler struct {
	store *EventStore
}

func NewCommandHandler(store *EventStore) *CommandHandler {
	return &CommandHandler{store: store}
}

func nowBase(t string) BaseEvent {
	return BaseEvent{Type: t, Timestamp: time.Now().UTC()}
}

func (h *CommandHandler) Handle(cmd Command) error {
	switch c := cmd.(type) {
	case OpenAccountCmd:
		// check account not exists
		if evs := h.store.Load(c.AccountID); len(evs) > 0 {
			return fmt.Errorf("account %s already exists", c.AccountID)
		}
		ev := AccountOpened{
			BaseEvent: nowBase("AccountOpened"),
			AccountID: c.AccountID,
			Owner:     c.Owner,
			Balance:   c.Initial,
		}
		h.store.Append(c.AccountID, ev)
		return nil
	case DepositCmd:
		evs := h.store.Load(c.AccountID)
		if len(evs) == 0 {
			return fmt.Errorf("account %s not found", c.AccountID)
		}
		if c.Amount <= 0 {
			return errors.New("amount must be positive")
		}
		ev := MoneyDeposited{
			BaseEvent: nowBase("MoneyDeposited"),
			AccountID: c.AccountID,
			Amount:    c.Amount,
		}
		h.store.Append(c.AccountID, ev)
		return nil
	case WithdrawCmd:
		evs := h.store.Load(c.AccountID)
		if len(evs) == 0 {
			return fmt.Errorf("account %s not found", c.AccountID)
		}
		// rehydrate and validate
		acc, err := RehydrateAccount(evs)
		if err != nil {
			return err
		}
		if c.Amount <= 0 {
			return errors.New("amount must be positive")
		}
		if acc.Balance < c.Amount {
			return errors.New("insufficient funds")
		}
		ev := MoneyWithdrawn{
			BaseEvent: nowBase("MoneyWithdrawn"),
			AccountID: c.AccountID,
			Amount:    c.Amount,
		}
		h.store.Append(c.AccountID, ev)
		return nil
	default:
		return fmt.Errorf("unknown command type %T", cmd)
	}
}

// -------------------- Projector / Read Model --------------------

// Read model: simple in-memory projection of account balances
type AccountView struct {
	AccountID string  `json:"account_id"`
	Owner     string  `json:"owner"`
	Balance   float64 `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Projector struct {
	store     *EventStore
	viewsLock sync.RWMutex
	views     map[string]AccountView
}

func NewProjector(store *EventStore) *Projector {
	return &Projector{
		store: store,
		views: make(map[string]AccountView),
	}
}

// rebuild view from scratch by reading all events (simple)
func (p *Projector) RebuildAll() {
	p.viewsLock.Lock()
	defer p.viewsLock.Unlock()
	p.views = make(map[string]AccountView)

	// iterate streams
	for streamID, evs := range p.store.events {
		acc := &Account{}
		for _, e := range evs {
			_ = acc.Apply(e)
		}
		p.views[streamID] = AccountView{
			AccountID: streamID,
			Owner:     acc.Owner,
			Balance:   acc.Balance,
			UpdatedAt: time.Now().UTC(),
		}
	}
}

// simple incremental projection: apply new events for a stream
func (p *Projector) ProjectStream(streamID string) {
	evs := p.store.Load(streamID)
	acc := &Account{}
	for _, e := range evs {
		_ = acc.Apply(e)
	}
	p.viewsLock.Lock()
	defer p.viewsLock.Unlock()
	p.views[streamID] = AccountView{
		AccountID: streamID,
		Owner:     acc.Owner,
		Balance:   acc.Balance,
		UpdatedAt: time.Now().UTC(),
	}
}

func (p *Projector) GetView(accountID string) (AccountView, bool) {
	p.viewsLock.RLock()
	defer p.viewsLock.RUnlock()
	v, ok := p.views[accountID]
	return v, ok
}

// Dump views to JSON (helper)
func (p *Projector) Dump() string {
	p.viewsLock.RLock()
	defer p.viewsLock.RUnlock()
	b, _ := json.MarshalIndent(p.views, "", "  ")
	return string(b)
}

// -------------------- Main (demo) --------------------

func main() {
	store := NewEventStore()
	handler := NewCommandHandler(store)
	projector := NewProjector(store)

	// 1. Open account
	if err := handler.Handle(OpenAccountCmd{AccountID: "acc-1", Owner: "Ton", Initial: 100}); err != nil {
		fmt.Println("err:", err)
		return
	}
	// project stream to update read model
	projector.ProjectStream("acc-1")

	// 2. Deposit
	_ = handler.Handle(DepositCmd{AccountID: "acc-1", Amount: 50})
	projector.ProjectStream("acc-1")

	// 3. Withdraw
	if err := handler.Handle(WithdrawCmd{AccountID: "acc-1", Amount: 30}); err != nil {
		fmt.Println("withdraw err:", err)
	}
	projector.ProjectStream("acc-1")

	// Query read model
	if v, ok := projector.GetView("acc-1"); ok {
		fmt.Printf("Account View: %+v\n", v)
	}

	// Show raw events
	fmt.Println("\nRaw events (per stream):")
	for sid, evs := range store.events {
		fmt.Printf("stream: %s\n", sid)
		for i, e := range evs {
			b, _ := json.Marshal(e)
			fmt.Printf(" %d: %s\n", i+1, string(b))
		}
	}

	// Simulate rebuild: new projector rebuilds projection from event store
	fmt.Println("\n-- Rebuild projection from events --")
	newProj := NewProjector(store)
	newProj.RebuildAll()
	fmt.Println(newProj.Dump())
}
