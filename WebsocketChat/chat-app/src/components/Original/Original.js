import React, { useState, useEffect } from "react";
function Original() {
  const [username, setUsername] = useState("");
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);
  const [ws, setWs] = useState(null);
  const [sockets] = useState([
    "ws://192.168.16.55:8080/ws",
    "ws://100.84.223.32:8080/ws",
  ]);
  const [isConnected, setIsconnected] = useState(false);

  useEffect(() => {
    // const socket = new WebSocket("ws://localhost:8080/ws");
    setUsername(localStorage.getItem("user_name"));
    statbleSocket().then((validSocketUrl) => {
      if (validSocketUrl) {
        alert("valid on :", validSocketUrl);
        setIsconnected(true);
        const socket = new WebSocket(validSocketUrl);

        socket.onmessage = (event) => {
          const newMessage = JSON.parse(event.data);
          setMessages((prev) => [...prev, newMessage]);
        };

        setWs(socket);
      } else {
        console.error("Don't have any valid websocket!");
      }
    });

    return () => ws && ws.close();
  }, []);

  async function statbleSocket() {
    if (sockets) {
      for (let i = 0; i < sockets.length; i++) {
        const isAvailable = await checkWebSocket(sockets[i]);
        if (isAvailable) {
          return sockets[i];
        }
      }
      return null;
    } else {
      return null;
    }
  }

  function checkWebSocket(url) {
    return new Promise((resolve) => {
      const testSocket = new WebSocket(url);

      testSocket.onopen = () => {
        testSocket.close();
        resolve(true);
      };

      testSocket.onerror = () => {
        resolve(false);
      };
    });
  }

  const sendMessage = () => {
    localStorage.setItem("user_name", username);

    if (ws && message.trim() !== "" && username.trim() !== "") {
      ws.send(JSON.stringify({ username, content: message }));
      setMessage("");
    }
  };
  return (
    <>
      <h1 className="text-center mt-5 pt-5 fw-bold">WebSocket Chat</h1>
      {isConnected ? (
        <>
          <div
            className="container"
            style={{ marginTop: "20px", padding: "20px" }}
          >
            <h2>User Name: {username ? username : "Insert name"}</h2>
            <input
              className="form-control"
              type="text"
              placeholder={username ? username : "Tên của bạn"}
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <br />
            <input
              className="form-control"
              type="text"
              placeholder="Nhập tin nhắn..."
              value={message}
              onChange={(e) => setMessage(e.target.value)}
            />
            <button
              className="btn btn-lg btn-success m-2"
              onClick={sendMessage}
            >
              Gửi
            </button>
            <div>
              <h3>Tin nhắn:</h3>
              <ul className="bg-secondary text-white border border-danger row d-flex">
                {messages.map((msg, index) => (
                  <li
                    key={index}
                    className="m-3 row d-flex justify-content-center align-items-center"
                  >
                    <div className="col-2 bg-danger">
                      <strong>{msg.username}:</strong>{" "}
                    </div>
                    <div className="col">
                      <span className="d-flex bg-primary rounded-pill m-2 p-2">
                        {msg.content}
                      </span>
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </>
      ) : (
        <h1 className="text-center">Loading...</h1>
      )}
    </>
  );
}

export default Original;
