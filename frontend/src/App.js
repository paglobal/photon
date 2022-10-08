import { html, ref, createRef } from "promethium-js";
import "./app.css";
import { ipcInit } from "photon-lib-js";

const App = () => {
  const input = createRef();
  const output = createRef();

  const ipc = ipcInit("53174");

  ipc.on("open", () => {
    output.value.innerHTML += "Status: Connected\n";
  });
  ipc.on("message", (payload) => {
    output.value.innerHTML += "Server: " + payload.message + "\n";
    output.value.innerHTML += "Server: " + payload.id + "\n";
  });

  const send = () => {
    ipc.emit("message", {
      message: input.value.value,
    });
    input.value.value = "";
  };

  return () => html`<div class="app">
    <span>Hello Proton</span>
    <div class="ipc-test">
      <input type="text" id="input" ${ref(input)} />
      <button @click=${send}>Send</button>
      <pre id="output" ${ref(output)}></pre>
    </div>
  </div> `;
};

export default App;
