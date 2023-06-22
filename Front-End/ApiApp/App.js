import './App.css';
import React, { useEffect, useState } from 'react';

function App() {

  const [lender, setLName] = useState("");
  const [receiver, setRName] = useState("");
  const [date, setDate] = useState("");
  const [amt, setAmt] = useState("");

  function saveUser() {
    console.warn({ lender, receiver, date, amt });
    let data = { lender, receiver, date, amt: parseInt(amt) };
    fetch("http://localhost:9090/transactions", {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }).then((result) => {
      result.json().then((resp) => {
        console.warn("resp", resp)
      })
    })
  }

  return (
    <div className="App">
      <h1>POST API EXAMPLE</h1>
      <div className="input-container">
        <input type="text" value={lender} onChange={(e) => { setLName(e.target.value) }} name="lender" placeholder="Lender" />
      </div>
      <div className="input-container">
        <input type="text" value={receiver} onChange={(e) => { setRName(e.target.value) }} name="receiver" placeholder="Receiver" />
      </div>
      <div className="input-container">
        <input type="text" value={date} onChange={(e) => { setDate(e.target.value) }} name="date" placeholder="Date" />
      </div>
      <div className="input-container">
        <input type="number" value={amt} onChange={(e) => { setAmt(e.target.value) }} name="amt" placeholder="    Amount" />
      </div>
      <button type="button" onClick={saveUser}>Save Data</button>
    </div>
  );

}

export default App;