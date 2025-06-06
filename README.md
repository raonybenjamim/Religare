Project Religare
=============

Religare is an instrumental trans communication application that aims to implement the Religare Protocol using golang and wifi signals as the modulation source.

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
- **Text Input**: Simply replicates data from a file at [./models/sampleTextContent](./models/sampleTextContent) (This is only for testing purposes)

## How to run? 

This application **should** run propperly on Windows and Linux systems. You can run the application with: 

```
go run main.go -config runConfigurations/desiredConfig.json
```

That will start the application and prompt you for the run configuration. Follow the on screen instructions to use the application.

To stop the application, use the `ctrl + c` command.

### Using the binaries

Each release also provide bundled binaries compatible with both windows and linux. You can use them instead of the `go run main.go` command. 

For example, if you are on windows, download the correct binary and click on it. 

> Note: Since we are not able to provide the correct certificates, the windows binaries may be stoped by windows security applications or your anti virus. If that happens, add the application to the exclusion list of our security applications or download the source code and run directly.

## Want to know more? 

You can access the full protocol definition [here (Portuguese text)](https://gentle-aura-fd4.notion.site/Proposta-Protocolo-Religare-ba51bc05f87542179d4187354ae60afd?pvs=74)

You can also reach me out at my [Instagram](https://www.instagram.com/espiritismoarretado) of [TikTok](https://www.tiktok.com/@raonybenjamim) profiles

## License 

This project is licensed under the GPL 3.0 license - see the LICENSE.md file for details.