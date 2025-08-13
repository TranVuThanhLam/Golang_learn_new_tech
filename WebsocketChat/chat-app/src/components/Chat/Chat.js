import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-icons/font/bootstrap-icons.css"; // Import Bootstrap icons
import { useNavigate } from "react-router-dom"; // Import useNavigate for navigation
import "./Chat.css"; // Additional custom styles

function Chat() {
  // websocket component
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
    const trimmedMessage = message.trim();
    const trimmedUsername = "Thanh Lam";
    // const trimmedUsername = username.trim();

    // Lưu username vào localStorage
    localStorage.setItem("user_name", trimmedUsername);

    // Nếu có nội dung hợp lệ và kết nối WebSocket
    if (ws && trimmedMessage !== "" && trimmedUsername !== "") {
      // Gửi qua WebSocket
      ws.send(
        JSON.stringify({ username: trimmedUsername, content: trimmedMessage })
      );

      // Thêm vào state messages để hiển thị
      setMessages((prev) => [
        ...prev,
        { text: trimmedMessage, sender: "user" },
      ]);

      // Xoá ô nhập
      setMessage("");
    }
  };

  // const sendMessage = () => {
  //   if (ws && message.trim() !== "" && username.trim() !== "") {
  //     ws.send(JSON.stringify({ username, content: message }));
  //     setMessage("");
  //   }
  // };

  // const handleSendMessage = () => {
  //   if (input.trim()) {
  //     setMessages([...messages, { text: input, sender: "user" }]);
  //     setInput("");
  //   }
  // };

  // Chat component

  const [input, setInput] = useState("");
  const [isSidebarOpen, setIsSidebarOpen] = useState(true); // State to toggle sidebar
  const navigate = useNavigate(); // Initialize navigate function

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };

  return (
    <div className="vh-100 d-flex">
      <div
        className={`chat-sidebar bg-dark text-white p-2 ${
          isSidebarOpen ? "col-2" : "col-1" // Use col-2 when expanded
        }`}
      >
        <div className="px-3">
          <button
            className="btn btn-dark text-white d-flex align-items-center mb-3 fw-bold fs-5 rounded-pill"
            onClick={() => navigate(-1)} // Navigate to the previous page
          >
            <i className="bi bi-arrow-left me-2 fs-4"></i> Back
          </button>
          <button
            className="btn btn-light d-flex align-items-center justify-content-center mb-3 rounded-pill"
            onClick={toggleSidebar}
          >
            <i className="bi bi-list fs-4"></i>{" "}
            {/* Updated icon to "menu" style */}
          </button>
        </div>
        {isSidebarOpen && (
          <>
            <h5 className="fw-bold text-center mt-5">People</h5>
            <hr />
            <div className="bg-secondary rounded">
              <ul>
                <li className="active-user p-1 rounded bg-white text-dark">
                  User 1
                </li>
                <li className="p-1">User 2</li>
                <li className="p-1">User 3</li>
              </ul>
            </div>
          </>
        )}
      </div>
      <div className="chat-main d-flex flex-column flex-grow-1">
        <div className="chat-header bg-black text-white text-center py-2 fw-bold">
          <h2>Chat - table booker</h2>
          <p className="mb-0">Chatting with: User 1</p>
        </div>
        <div className="chat-messages flex-grow-1 p-2 overflow-auto d-flex flex-column">
          {messages.map((message, index) => (
            <div
              key={index}
              className={`chat-message rounded-pill px-3 py-1 mb-1 fs-6 ${
                message.sender === "user"
                  ? "bg-primary text-white align-self-end" // Sent messages on the right
                  : "bg-danger text-white align-self-start" // Received messages on the left
              }`}
            >
              {message.text}
            </div>
          ))}
        </div>
        <footer className="chat-footer bg-dark text-white d-flex p-2">
          <input
            type="text"
            className="form-control me-1 rounded-pill px-2"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            placeholder="Type a message..."
            onKeyDown={(e) => e.key === "Enter" && sendMessage()}
          />
          <button
            className="btn bg-success rounded-pill d-flex align-items-center mx-2"
            onClick={sendMessage}
          >
            <i className="bi bi-send me-1 px-2"></i> {/* Added send icon */}
          </button>
        </footer>
      </div>
    </div>
  );
}

export default Chat;
