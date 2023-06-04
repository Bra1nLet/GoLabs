import React from 'react';
import { useRef, useState } from 'react';
import {Authorize}  from "../../hooks";
import {ExitFromAccount, GetToken} from "../../hooks/Auth";
import './Drive.css'

const Drive = () => {
    const [filesAll, setFilesAll] = useState(null);
    const [folderssAll, setFoldersAll] = useState(null);
    const [loaded, setLoaded] = useState(false);
    const path = useRef("");
    const handleButtonClick = () => {
        const formData = new FormData();
        formData.append('path', path.current.value);
        const authOptions = {
            method: "POST",
            body: formData,
            headers: {
                'Authorization': GetToken()
            },
        };

        fetch('/drive/get-folder', authOptions)
            .then(response => response.json())
            .then(response => {setFilesAll(response.files); setFoldersAll(response.folders)})
            .then(response => setLoaded(true))
            .catch(error => {
                console.error('Error:', error);
            });
    };



    return (
        <div>
            <button onClick={ExitFromAccount}>Logout</button>
            <br/>
            <br/>
            {loaded ?
                <div>
                    <h1>List of Files:</h1>
                    <ul>
                        {folderssAll.map((item, index) => (
                            <li className="folder" key={index}>{item}</li>
                        ))}
                        {filesAll.map((item, index) => (
                            <li className="file" key={index}>{item}</li>
                        ))}
                    </ul>
                </div> : handleButtonClick()
            }
            <label htmlFor="path">Path</label>
            <br />
            <input ref={path} name="path" type="text" />
            <br />
            <button onClick={handleButtonClick}>In directory</button>
        </div>
    );
};
export default Drive;