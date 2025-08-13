import React from "react";
import Original from "./components/Original/Original";
import Chat from "./components/Chat/Chat";
import { BrowserRouter } from "react-router-dom";
const App = () => {
  return (
    <>
      <BrowserRouter>
        <Chat />
        {/* <Original /> */}
      </BrowserRouter>
    </>
  );
};

export default App;
