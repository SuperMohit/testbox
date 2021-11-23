<h1 align="center">Welcome to TestBox 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-v1-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> This project consists of 3 services. 
> 1. Signup service: This services Signs up the user, creates a subdomain using Route 53 API and then Sends an asynchronous request to Orchestrator Service. 
> 2. Orchestrator service: This service applies docker compose template to schedule a docker which listens at the random port in Ec2 machine. 
> 3. Service-server: This server provides interface to the user using separate docker instances based on domain, by which user can see BB, have access to isolated Databse and upload to s3 bucket.

## Author

* Website: talniya.com
* Github: [@supermohit](https://github.com/supermohit)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
