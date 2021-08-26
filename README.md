# multi-store-management-api

The project is an API that represents a simplified backend for a multi-tenant e-commerce platform (like shopingify).<br/>
This project is simply a proof of concept that I am using to learn Golang and it is not tested enough for a production level project, as it has been developed in about two weeks with absolutely no prior knowledge of Go.<br/>
We are using Gin for the API and the official MongoDB driver for communicating with the database.<br/>

## Running the project
Inside the project directory run the command:<br/>
`go run server.go`<br/>
You should see the message *"Listening and serving HTTP on localhost:8080"* at the end.<br/>

## Project Structure
For this project, the modules where defined based on business requirements.<br/>
* meta: This module is responsible for managing the meta-data, which in this case concerns the tenant's information that is not considered as "content" (an example is the title they chose for the website).<br/>
* menu-items: This module is responsible for managing the items listed in the website's main menu as they can vary from one tenant to another.
* product-lists: This module is responsible for the custom lists that the tenant might choose to display in their home page.
* products: This module is responsible for managing the products as the main content of the application.
* shared: This is a directory regrouping multiple modules that are used accross the project.
  * database: a module that offers function to facilitate manipulating MongoDb collections as well as the connection to the database.
  * entities: a module containing all database entities.
  * genericRepository: a module providing the basic repository functionalities (find, create, update, delete, etc...) to avoid re-writing them for each entity.
  * models: a module containing structures used in multiple modules of the project. Unlike the structures in *entities*, the ones in *models* are not persisted in the database.

## Technical choices
In multi-tenant applications, there is usually a choice between 3 possibilities:
* Placing all data their respective collections and adding a field to all schemas to indicate which tenant they belong to.
* Using different collections for each tenant
* Using different databases for each tenant

Each of these solutions present their own advantages and disaventages.<br/>
The first one is the simplest to implement but is the heaviest in terms of query processing, as a search for a given product for example, will be conducted on the products of all tenants instead of just the one who's Hostname is being used. This obviously get worse the more tenants we have. So, unless we are counting on a very limited number of tenants, each having a relatively small number of database entries, this solutions is to be dismissed.<br/>
The second solution is a bit harder to impliment compare to the first one but provides much more scalability. However, the MongoDB official documentation recommends not exceeding 10,000 collections per replica-set. This puts a limit on this solution's scalability, but is still much more reliable than the first solution.<br/>
The thirst solution is the most scalable of the 3, but also the most costly as it uses multiple databases which, when not justified but having large volumes of data for the tenants, can cause a counter productive overhead cost. This is due to the fact that the backend must support multiple connections to multiple databases, as well as the overhead of running multiple databases on a clusted as opposed to one.<br/><br/>

Taking into consideration all these factors, I decided to use the second approach which can later be upgraded to a hybrid approach where multiple databases are used, but each database is used to store the data of a group of tenants instead of just one. This allows for more flexibility and more importantly for not over-engineering the solution with an expensive architecture that is not yet needed or justified.<br/>

To implement the multi-collection multi-tenant solution, I used a configuration file that links the hostname used in the query to a collection name prefix which allows to identify which collection should be used for the request currently being handled. For the moment, this file must be updated manually for the sake of simplicity but it can be upgraded to a combination of a backoffice to manage new entries, a database to persist those entries and an in-memory data store (such as Redis) to enhance performance by avoiding to query the database for each request before even starting to run the business logic code.
