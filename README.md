Project Religare
=============

Religare is an instrumental trans communication that aims to implement the Religare Protocol using golang and wifi signals as the modulation source.

## What is the Religre Protocol?

The communication protocol aims to facilitate communication with entities present in the spiritual plane and any entity capable of manipulating the electromagnetic spectrum or any other form of signal that can be captured and converted to binary format. This protocol allows the construction of a direct communication channel with these entities, bypassing the limitations of human mediumship.

## What to expect from this application?

Through a set of definitions, this protocol aims to create a standardized way of communicating directly with entities capable of manipulating the electromagnetic spectrum or other forms of communication signals, bypassing the limitations of human mediumship.

We assume the ability of these entities to influence these signals and communicate through this influence.

This protocol envisions the possibility of transmitting any binary message. This includes, but is not limited to:

- Text messages
- Images
- Audios
- Videos

### Disclaimer
The implementation of this protocol is for experimental purposes only. The project does not confirm the moral elevation of communicating entities and advises caution in message evaluation.

### Features
- **Direct Communication**: Establishes a direct channel with entities without the limitations of human mediumship.
- **Binary Transmission**: Supports transmission of binary messages including text, images, audio, and video.

### Components
- **Communication Channel**: Converts various signals into binary data.
- **Converter**: Transforms signal variations into binary values based on a defined threshold.
- **Interpreter**: Encodes binary data into the final message format.

### Running modes: 
- **wifi**: Reads binary data based on the oscilation of the WiFi signal strenght
- **Random**: Reads randomly generated binary data based on an internal algorithm
- **Text Input**: Simply replicates data from the console input (This is only for testing purposes)

## How to run? 

You can run the application with: 

```
go run main.go
```

That will start the application in the WiFi mode. If you want to run in a different mode, provide the `generator-type` parameter: 

```
go run main.go --generator-type=TextInput
```