import React, { useState } from 'react';
import axios from 'axios';
import { QRCode } from 'react-qrcode-logo';
import Img from './image/logo.png'

function App() {
  const [url, setUrl] = useState('');
  const [qrCodeId, setQrCodeId] = useState(null);
  const [error, setError] = useState(null);

  const handleGenerateQRCode = async () => {
    try {
      setError(null); // Reset any previous errors
      const response = await axios.post('http://localhost:8080/generate', { url });
      setQrCodeId(response.data.id);
    } catch (err) {
      setError('Failed to generate QR code. Please try again.');
    }
  };

  return (
    <div style={{ textAlign: 'center', marginTop: '50px' }}>
      <h1>QR Code Generator</h1>
      <input
        type="text"
        placeholder="Enter URL"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        style={{ padding: '10px', width: '300px', marginBottom: '20px' }}
      />
      <br />
      <button onClick={handleGenerateQRCode} style={{ padding: '10px 20px' }}>
        Generate QR Code
      </button>
      <div style={{ marginTop: '20px' }}>
        {error && <p style={{ color: 'red' }}>{error}</p>}
        {qrCodeId && (
          <>
            <QRCode value={url} size={256} logoImage={Img} logoPadding={3} ecLevel="M" bgColor="white" fgColor="black"
              qrStyle='dots'
            />
            <p>
              QR Code ID: {qrCodeId} <br />
              Scan this QR code or <a href={`http://localhost:8080/qr/${qrCodeId}`}>click here</a> to open the URL.
            </p>
          </>
        )}
      </div>
    </div>
  );
}

export default App;
