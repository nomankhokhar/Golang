// require('@grpc/grpc-js') imports the core gRPC library for Node.js to handle connections, streaming, and RPC communication.
const grpc = require('@grpc/grpc-js');

// require('@grpc/proto-loader') imports the utility to dynamic-load and parse .proto files at runtime.
const protoLoader = require('@grpc/proto-loader');

// require('path') imports Node's built-in path module to navigate directories reliably across different platforms.
const path = require('path');

// PROTO_PATH defines the path to the 'pingpong.proto' file inside the Golang directory.
// - __dirname is the directory of the current file (GRPC_Practice/NodeJS/client).
// - '..', '..', 'Golang', 'pingpong.proto' traverses up to the shared proto definition file.
const PROTO_PATH = path.join(__dirname, '..', '..', 'Golang', 'pingpong.proto');

// protoLoader.loadSync reads and parses the .proto file synchronously, producing a package schema definition object.
const packageDefinition = protoLoader.loadSync(PROTO_PATH);

// grpc.loadPackageDefinition loads the parsed schema into gRPC-usable Javascript namespaces and client classes.
// - .pingpong targets the specific package name defined inside the proto file.
const proto = grpc.loadPackageDefinition(packageDefinition).pingpong;

// Create a gRPC client instance of GameService.
// - 'localhost:50051': Points to the Golang gRPC server address and port.
// - grpc.credentials.createInsecure(): Establishes an unencrypted channel (no SSL/TLS) for local testing.
const client = new proto.GameService('localhost:50051', grpc.credentials.createInsecure());

// client.PlayPingPong() initiates a bidirectional stream call to the server.
// - call represents the client stream connection (it implements both Writable stream to send data, and Readable stream to receive data).
const call = client.PlayPingPong();

// call.on('data') is an event listener that fires whenever a new message is received from the server.
call.on('data', (res) => {
    // res is the response message object from the server. We check if res.text is 'Ping'; if so, we reply with 'Pong', otherwise 'Ping'.
    const reply = res.text === 'Ping' ? 'Pong' : 'Ping';
    
    // Log the incoming message and our planned reply to the console.
    console.log(`${res.text} → ${reply}`);
    
    // setTimeout delays sending the reply back by 1000 milliseconds (1 second) to keep the ping-pong pace readable.
    // - call.write sends the reply object ({ text: reply }) through the stream back to the server.
    setTimeout(() => call.write({ text: reply }), 1000);
});

// call.on('error') registers a listener to print out any connection errors or stream crashes.
call.on('error', (err) => console.error(err.message));

// call.write sends the initial payload to boot up the game.
// - We send '{ text: 'Ping' }' to trigger the first 'data' event on the Golang server.
call.write({ text: 'Ping' });
