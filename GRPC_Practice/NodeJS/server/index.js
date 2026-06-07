// require('@grpc/grpc-js') imports the core gRPC implementation for Node.js.
// This handles network connectivity, serialization, streams, and low-level HTTP/2 protocols.
const grpc = require('@grpc/grpc-js');

// require('@grpc/proto-loader') imports a utility package to load .proto definition files.
// It parses the protobuf file dynamically, converting it into a structure that can be loaded by @grpc/grpc-js.
const protoLoader = require('@grpc/proto-loader');

// require('path') imports the Node.js native path module to resolve and manipulate file/directory paths.
const path = require('path');

// PROTO_PATH calculates the absolute path to the 'pingpong.proto' file.
// - __dirname is the directory of the current file (GRPC_Practice/NodeJS/server).
// - '..', '..', 'Golang', 'pingpong.proto' navigates up two folders and into the Golang folder to locate the proto file.
const PROTO_PATH = path.join(__dirname, '..', '..', 'Golang', 'pingpong.proto');

// protoLoader.loadSync reads the proto file synchronously and parses it into a package definition schema.
// This is necessary because JavaScript is dynamic and compiles the API structure at runtime rather than compile-time.
const packageDefinition = protoLoader.loadSync(PROTO_PATH);

// grpc.loadPackageDefinition takes the schema definitions and generates service & message definitions.
// - .pingpong accesses the package namespace declared inside the pingpong.proto file ('package pingpong;').
const proto = grpc.loadPackageDefinition(packageDefinition).pingpong;

// playPingPong is our implementation of the GameService.PlayPingPong bidirectional streaming handler.
// - call represents the active streaming context. It acts as both a Readable stream (for client messages) and a Writable stream (for server responses).
function playPingPong(call) {
    // call.on('data') is an event listener that fires whenever a new message is received from the client stream.
    call.on('data', (req) => {
        // req is the received message object. We check if req.text is 'Ping'; if so, we respond with 'Pong', otherwise 'Ping'.
        const reply = req.text === 'Ping' ? 'Pong' : 'Ping';
        
        // Log the received message and the server response to the console.
        console.log(`${req.text} → ${reply}`);
        
        // call.write sends the response message back to the client. The object format matches the structure defined in the proto (message Message { string text = 1; }).
        call.write({ text: reply });
    });
    
    // call.on('end') is an event listener that fires when the client closes their end of the stream.
    // When the client is done, we close the server side of the stream too by calling call.end().
    call.on('end', () => call.end());
}

// Create a new gRPC Server instance to accept incoming RPC connections.
const server = new grpc.Server();

// server.addService registers our playPingPong handler to the proto's GameService definitions.
// - proto.GameService.service is the generated service structure specifying methods and types.
// - { playPingPong } is an object matching the proto method name (PlayPingPong, camel-cased by loadSync to playPingPong) to our JS handler.
server.addService(proto.GameService.service, { playPingPong });

// server.bindAsync binds the server to a network interface and port.
// - '0.0.0.0:50052' binds to all network interfaces on port 50052.
// - grpc.ServerCredentials.createInsecure() configures the server to run in cleartext mode (no SSL/TLS/auth).
// - The final callback function executes once binding succeeds.
server.bindAsync('0.0.0.0:50052', grpc.ServerCredentials.createInsecure(), () => {
    // Confirm to the console that the server has started and is listening.
    console.log('Node.js gRPC Server listening on :50052');
});
