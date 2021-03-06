## Open climate

This is the first repository for the **open climate project**, an open source initiative exploring the application of **distributed ledger technology (DLT)** and other emerging technologies, such as **IoT (Internet of Things), big data and machine learning**, to the challenge of helping the world keep a transparent climate accounting system towards the climate targets set in the 2015 **Paris Agreement—i.e. maintaining anthropogenic warming below 1.5oC**.

The repo is used for the first demo versions of the platform, which explore the initial scope of the system and mocks up functionalities. 

This project involves a software ‘platform of platforms,’ distinguished here as the Open Climate platform (capitalized to distinguish from the overall open climate project), the development of climate communication protocols, and a shared user interface as portal to the system. **Front end repository fro the demo can be found [here](https://github.com/YaleOpenLab/openx-frontend)**. The platform is designed to ideally acts as a common integrator designed to reconcile climate records and functions from both legacy and DLT-based climate platforms in the pursuit of helping maintain a decentralized ‘ledger of ledgers’.

Whilst the platform is currently incubated at the Yale Open Innovation Lab (openlab) it seeks to combine a growing network of other platforms in advanced technological stages; using and open innovation effort to develop the common open source layers. 

For more information, visit: [https://openlab.yale.edu/open-climate](https://openlab.yale.edu/open-climate).

## Prerequisites and installation

Clone the repository to a directory of your choosing:

``git clone https://github.com/YaleOpenLab/openclimate-demo.git``

Install latest distribution of [Go](https://golang.org/) if you don't have it already.

Finally, build the program using the following command.

```
cd openclimate-demo 
go build
```

Run the program.

``./openclimate``

For blockchain smart contract environment please refer to this [instructions](https://github.com/YaleOpenLab/openclimate-demo/blob/master/blockchain/README.md)
