package main

// what is the consumer key?

// An API key (known as a consumer key) is a string value passed by a client app to your API proxies.
// The key uniquely identifies the client app.
//  API key validation is the simplest form of app-based security that you can configure for an API.

// what is the access token ?

// User login base token authentication methods:

// Here are four ways to get into applications:
//  The common goal of these four methods is to gain
//  access token and use it to access information and authentication.

//  Authorization Code Grant: A string of code used to issue an access token.
// This code is published when the user logs in to Front and is created in front of the access token on the server side.
//  As a result, the user is authenticated with the password and code generated.

//  Implicit Grant: When the user logs in, the corresponding access token is created.

// Client Credential Grant: Access token is created on the server side only for client (not user) authentication.

// Password (Grant): access token immediately with a single request containing user login
//  information (username and password or client id and client secret) This method is easier to implement and also has drawbacks.

// Authorization Code Grant :
// Login steps are done in two steps and with a high level of security.

// Keywords and definitions:

// Client App:
// In this way, the client is generally a web application and must include front and back.
// This means that an empty site with Angular or React or ... can not run this method.
// (Of course, it can use the Implicit Grant method, which will be mentioned below).

// Authorization Server:
// The part that performs authentication and authorization
// (ie user login requests and user authentication and creating tokens and security validations is the responsibility of this section) .

// Resource Server:
//  The part that makes information available, etc., and this can be the REST API.
// After receiving the access token by Front, it will be accessible and usable.

// One of the differences between an authentication server and a resource server is the type of work
//  they do (the authentication server handles only login and token requests, etc., and the resource server handles only content and information).

// Bearer Tokens are the predominant type of access token used with OAuth 2.0.
// A Bearer Token is a string, not intended to have any meaning to clients using it.
// Some servers will issue tokens that are a short string of hexadecimal characters,
//  while others may use structured tokens such as JSON Web Tokens
