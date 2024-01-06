# design url shortener

### 1. Requirements for System Design of URL Shortner Service

#### 1.1 Functional requirements of URL Shortening service

- Given a long URL, the service should generate a shorter and unique alias for it.
- When the user hits a short link, the service should redirect to the original link.
- Links will expire after a standard default time span.

#### 1.2 Non-Functional requirements URL Shortening service

- The system should be highly available. This is really important to consider because if the service goes down, all the URL redirection will start failing.
- URL redirection should happen in real-time with minimal latency.
- Shortened links should not be predictable.

**Note**: Let’s consider we are using 7 characters to generate a short URL. These characters are a combination of 62 characters [A-Z, a-z, 0-9] something like http://ad.com/abXdef2.