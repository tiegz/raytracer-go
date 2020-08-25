importScripts("wasm_exec.js");

if (!WebAssembly.instantiateStreaming) { // polyfill
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}

const go = new Go();
let mod, inst;

// Handle incoming messages
self.addEventListener('message', async function (event) {
  const { eventType, eventData } = event.data;
  // console.log("WORKER", eventType, eventData)
  switch (eventType) {
  case "INITIALISE":
    WebAssembly.instantiateStreaming(fetch("raytracer.wasm"), go.importObject).then((result) => {
      mod = result.module;
      inst = result.instance;
      go.run(inst)
    }).then((result) => {
      // inst = WebAssembly.instantiate(mod, go.importObject); // reset instance
      self.postMessage({ eventType: "INITIALISED", eventData: null });
    }).catch((err) => {
      console.error(err);
    });
    break;
  case "RENDER":
    self.postMessage({
      eventType: "RENDERED",
      eventData: {
        x: eventData[0],
        y: eventData[1],
        color: renderPixel(eventData[0], eventData[1])
      }
    });
    break;
  default:
    console.log("Unknown event: ", eventType, eventData)
    break;
  }
}, false);
