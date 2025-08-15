# Builder pattern

The Builder pattern is used when the desired product is complex and requires multiple steps to complete. In this case, several construction methods would be simpler than a single monstrous constructor. The potential problem with the multistage building process is that a partially built and unstable product may be exposed to the client. The Builder pattern keeps the product private until it’s fully built.

<p align="center">
  <img src="https://media.geeksforgeeks.org/wp-content/uploads/uml-of-builedr.jpg">
</p>