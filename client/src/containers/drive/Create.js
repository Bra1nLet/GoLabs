import React from 'react';
import { useRef, useState } from 'react';
import {Authorize}  from "../../hooks";
import {ExitFromAccount, GetToken} from "../../hooks/Auth";
import './Drive.css'

const CreateDrive = () => {
    const path = useRef("");
    const pathFolder = useRef("");
    const nameFile = useRef("");

    const [selectedFile, setSelectedFile] = useState(null);

    const handleFileChange = (event) => {
        setSelectedFile(event.target.files[0]);
    };

    const handleButtonCreate = () => {
        const formData = new FormData();
        formData.append('path', pathFolder.current.value);
        formData.append('folderName', nameFile.current.value);
        const authOptions = {
            method: "POST",
            body: formData,
            headers: {
                'Authorization': GetToken()
            },
        };

        fetch('/drive/new-folder', authOptions);
        //window.location.reload(false);
    };

    const handleButtonUpload = () => {
        const formData = new FormData();
        formData.append('path', path.current.value);
        formData.append('file', selectedFile);
        const authOptions = {
            method: "POST",
            body: formData,
            headers: {
                'Authorization': GetToken()
            },
        };

        fetch('/drive/upload', authOptions);
        //window.location.reload(false);
    };


    return (
        <div>
            <h1>Create Directory</h1>
            <label htmlFor="path">Path</label>
            <br />
            <input ref={pathFolder} name="path" type="text" />
            <br />
            <label htmlFor="name">Name</label>
            <br />
            <input ref={nameFile} name="name" type="text" />
            <br />
            <button onClick={handleButtonCreate}>Create new folder</button>
            <br />
            <br />
            <h1>Upload File</h1>
            <label htmlFor="path">Path</label>
            <br />
            <input ref={path} name="path" type="text" />
            <br />
            <br />
            <input onChange={handleFileChange} type="file"></input>
            <button onClick={handleButtonUpload}>Upload file</button>
        </div>
    );
};
export default CreateDrive;