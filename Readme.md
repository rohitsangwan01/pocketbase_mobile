# Pocketbase Mobile

Pocketbase mobile is used to generate android and ios packages for using pocketbase in mobiles


## To build

Make sure [gomobile](https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile) is installed 

run :  `gomobile bind -androidapi 19` for Android, this will generate `pocketbaseMobile.aar`, import this in android and use

run : `gomobile bind --target ios` for IOS or try : `gomobile bind -ldflags='-extldflags=-libresolv.tbd' -target=ios`

# Usage

[pocketbase_server_flutter](https://github.com/rohitsangwan01/pocketbase_server_flutter) to use in Flutter

[pocketbase_android](https://github.com/rohitsangwan01/pocketbase_android) to use in native Android
 
[pocketbase_ios](https://github.com/rohitsangwan01/pocketbase_ios) to use in native IOS

# Examples


checkout [Pocketbase Server Flutter](https://github.com/rohitsangwan01/pocketbase_server_flutter) for android and ios implementation in flutter

<img src="https://github.com/rohitsangwan01/pocketbase_server_flutter/assets/59526499/7d20a2a4-0df7-4f2a-90bf-2577289e0f7e" height="300">
<img src="https://github.com/rohitsangwan01/pocketbase_server_flutter/assets/59526499/370c007d-51c3-45a9-928c-1287c8def0d3" height="300">
<img src="https://github.com/rohitsangwan01/pocketbase_server_flutter/assets/59526499/657a6e4c-8431-4f49-b29d-a0f599524f6c" height="300">
<img src="https://github.com/rohitsangwan01/pocketbase_server_flutter/assets/59526499/4ecd5f1c-ae2b-4406-a10d-0d9ae3e9900e" height="300">
<img src="https://github.com/rohitsangwan01/pocketbase_mobile/assets/59526499/5ec533af-1b6f-4c79-afd8-e3e65e2d55a1" height="300">
<img src="https://github.com/rohitsangwan01/pocketbase_server_flutter/assets/59526499/f58f7f5e-d3d0-4328-a8be-f5cf12e15cdb" height="300">

checkout [Pocketbase Server Android](https://github.com/rohitsangwan01/pocketbase_server_android_example) for native android implementation

<img src="https://github.com/rohitsangwan01/pocketbase_mobile/assets/59526499/ff2c277a-bc9e-456c-b089-42fd264f61e3" height="300">
<img src="https://github.com/rohitsangwan01/pocketbase_mobile/assets/59526499/93b668c8-600f-4232-b2bb-3562ccbde32e" height="300">

# Usecase

- To create a secure localChat server or maybe a file storage system or anything which provided by pocketBase
- To use an old android device as local server to host PocketBase
- To use a mobile app which is based on PocketBase, so we can simply run PocketBase server within mobile and show demo of our application, useful for someone who don't want to host pocketBase yet
- To use Pocketbase as a College project,and for students team who want to show demo of a mobile app but can't afford hosting or something , so all they have to do is , run pocketBase server from there mobile and their main app which is based on pocketBase is ready to use

# Extras

Checkout a Flutter chatApp built using pocketbase: [flutter_pocketbase_chat](https://github.com/rohitsangwan01/flutter_pocketbase_chat)


