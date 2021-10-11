# Jarvis-Personal-Linux-Assistant

[![Build Status](https://travis-ci.com/Harkishen-Singh/Jarvis-personal-assistant.svg?branch=master)](https://travis-ci.com/Harkishen-Singh/Jarvis-personal-assistant)

[![dfsdf](https://files.gitter.im/COSS-Jarvis/community/euO1/tumblr_nrqm32yH3W1r6xm5co1_1280.gif)](https://gitter.im/COSS-Jarvis/community)

## Introduction

The project aims to develop a personal-assistant for Linux-based systems. Jarvis draws its inspiration from virtual assistants like Cortana for Windows, and Siri for iOS. It has been designed to provide a user-friendly interface for carrying out a variety of tasks by employing certain well-defined commands. Users can interact with the assistant either through ***voice commands or using a keyboard input***.

## Update with the project
The project service has been revamped using NodeJS. To have a look on the project and start contributing visit the branch `node-master` branch in the same repository or [click here](https://github.com/Harkishen-Singh/Jarvis-personal-assistant/tree/node-master).

## Deployment
The backend has been deployed in Heroku. 

The link to the API for GoLang based service is:
```
https://assistant-jarvis.herokuapp.com/
```
The link to the API for Node based service:
```
https://assistant-jarvis-node.herokuapp.com/
```

## Getting Started

To know the steps to install and run the project see [INSTALL.md](https://github.com/Harkishen-Singh/Jarvis-personal-assistant/blob/master/INSTALL.md)

**Starting Jarvis in development mode**

Installation:

1. Install all dependencies: `make install-all`

Update dependencies: `make update`

Run Jarvis:

1. Run service: `make run`
2. Run the desktop app: `make views`

To get started with your contributions for Jarvis-personal-assistant see [CONTRBUTING.md](https://github.com/Harkishen-Singh/Jarvis-personal-assistant/blob/master/CONTRIBUTING.md)

### Explanatory Video
[![Jarvis explanatory video](https://files.gitter.im/Harkishen-Singh/QIzs/Screenshot-_38_.png)](https://youtu.be/jztI_iN82RY)

## What Jarvis can do?

As a personal assistant, Jarvis assists the end-user with day-to-day activities like *general human conversation, searching queries in google, bing or yahoo, searching for videos, retrieving images, live weather conditions, word meanings, searching for medicine details, health recommendations based on symptoms and reminding the user about the scheduled events and tasks*. The user statements/commands are analysed with the help of **machine learning** to give an optimal solution.

# Desktop Application

## Introduction

The project aims to develop a personal-assistant for Linux-based systems. Jarvis draws its inspiration from virtual assistants like Cortana for Windows, and Siri for iOS. It has been designed to provide a user-friendly interface for carrying out a variety of tasks by employing certain well-defined commands.

This is the desktop version of jarvis personal assistant. Click [here](http://github.com/Harkishen-Singh/Jarvis-Personal-Assistant) to view the repository of server.

[Installation Instructions](https://github.com/muskankhedia/Jarvis-Desktop/blob/master/INSTALL.md)

## :wrench: Techology stack
* **Frontend** AngularJS,ElectronJS

## :rocket: Features 

**Queries from web**<br/>
In order to make queries from different search engines, the given format should be adopted.
For making queries from google
Google your query<br/>
Similarly for yahoo and bing.<br/>
i.e. google APJ Abdul Kalam

**Accessing youtube videos**<br/>
In order to access videos from youtube format is,<br/>
Youtube “video you want to search for”<br/>
i.e. youtube Sandeep Maheshwari

**Get weather for a location**<br/>
To get the weather at any location format is,<br/>
Weather city state<br/>
i.e. weather Bhubaneswar Odisha

**Retrieve images**<br/>
For retrieving images format is,<br/>
Image “image you want to search”<br/>
i.e. image "M S Dhoni"

**Dictionary meaning**<br/>
For getting meaning of a word format is,<br/>
meaning "word you are searching for"<br/>
i.e. meaning ecstasy

**Get medicine details**</br>
For getting details about any specific medicine, format is<br/>
medicine "name of the medicine"<br/>
i.e. medicine paracetamol

**Retrack causes of symptoms**<br/>
For retracking causes of symptoms, format is<br/>
symptoms "disease/ailment"<br/>
i.e. symptoms headache

**Set Reminders**<br/>
In order to set reminder type<br/>
set reminder<br/>
and fill in the necessary details.<br/>
i.e. set reminder 

**Sending Emails**<br/>
In order to send email, format is<br/>
send email<br/>
and fill in the necessary details to send the email.<br/>
i.e. send email

**Deploy at heroku**<br/>
In order to deploy, the github repo should be provided with proper Procfile config,<br/>
deploy <name of repo><br/>
i.e. deploy move-hack-angular

**Medicine Details**<br/>
Get the complete details of the medicine, including :
1. Indications
2. Contradictions
3. Trade/Brand Names
4. Indications
5. Dosage
6. Process of consumption
7. Warnings and Precautions related to the medicine
8. Storage conditions

**Medicine Help based on symptoms**<br/>
List your noticed symptoms and get immediate help on the medicines that could be taken.

### DFD explaining the data flow in the software:<br>
![DFd expalining data flow](https://files.gitter.im/muskankhedia/inED/moqup-_1_.png)

## :wrench: Tech Stack

* **Front-end:**  Web: AngularJS, Desktop App: Electron + AngularJS
* **Back-end:** GoLang
## :computer: Postman Documentation
Postman is a powerful HTTP client for testing web services. Postman makes it easy to test, develop and document APIs by allowing users to quickly put together both simple and complex HTTP requests. <br/>
Test the current APIs of Jarvis mentioned here on Postman - [Postman Link](https://documenter.getpostman.com/view/6521254/SzKZrvTu?version=latest)

#  :trophy: Achievements
 ### Jarvis-personal-assistant has been selected many times in various coding programmes like 
  
 #### KWoC-2019 (Kharagpur Winter of Code)
[<img src="https://miro.medium.com/max/585/1*IyO7IeMmLgTe0sgwR0Koeg.jpeg" width="100" height="100" />](https://wiki.metakgp.org/w/Kharagpur_Winter_of_Code)


#### JMoC-2019 (JIIT Month of Code)
[<img src="https://jmoc.jodc.tech/logo.png"/>](https://jmoc.jodc.tech/)


#### GSSoC-2020 (GirlScript Summer of Code)
[<img src="https://raw.githubusercontent.com/GirlScriptSummerOfCode/MentorshipProgram/master/GSsoc%20Type%20Logo%20Black.png" width="400" height="100" />](https://www.gssoc.tech/index.html)
