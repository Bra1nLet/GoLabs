import React, { useRef } from 'react';
import { GetToken } from "../../hooks/Auth";
import './Drive.css';

const DownloadDrive = () => {
  const path = useRef("");
  const fileName = useRef("");

  const handleButtonCreate = async () => {
    const formData = new FormData();
    formData.append('path', path.current.value);
    formData.append('filename', fileName.current.value);
    const authOptions = {
      method: "POST",
      body: formData,
      headers: {
        'Authorization': GetToken()
      },
    };

    try {
      const response = await fetch('/drive/download', authOptions);
      const blob = await response.blob();
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = fileName.current.value;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
    } catch (error) {
      console.error("Error downloading file:", error);
    }
  };

  return (
    <div>
      <h1>Download file</h1>
      <label htmlFor="path">Path</label>
      <br />
      <input ref={path} name="path" type="text" />
      <br />
      <label htmlFor="name">Name</label>
      <br />
      <input ref={fileName} name="name" type="text" />
      <br />
      <button onClick={handleButtonCreate}>Download file</button>
    </div>
  );
};

export default DownloadDrive;
