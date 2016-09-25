# goa-design-base
The purpose of this repository is to seed a new goa and gorma design folder. 

### Steps to Seed a new Project

1) Perform a go get on this repository - note that it will not compile - that is fine
2) Update the utils/moeboe.properties with your desired names
3) Execute the ./new_goa.sh file and follow the directions.
4) See the [Sample Goa Project](https://github.com/cboecking/gorma-sample) to see an example project created from this seed.

### Steps to make your project work

1) Follow the directions from ./new_goa.sh to test the seeded project
2) Install a postgresql instance with the appropriate credentials. [Here is a docker command if you docker installed.](https://github.com/goadesign/gorma-cellar/blob/master/docker.sh)
3) Update main.go with your database information: see the [Sample Goa Project](https://github.com/cboecking/gorma-sample)
4) Update your controller (ex: account.go) to update the database based on the routing actions: see the [Sample Goa Project](https://github.com/cboecking/gorma-sample)

### Feedback
I hope this helps! Let me know if you have feedback: chuck@chuboe.com.
