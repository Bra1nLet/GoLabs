import logo from './logo.svg';
import './App.css';
import {Drive, CreateDrive, Download, Registration, DeleteDrive, Rename, Login} from "./containers";
import {GetUser, ValidateToken} from "./hooks/Auth";

function App() {
  if (!ValidateToken()){
    return (
        <div className="App">
          <Login/>
          <Registration/>
        </div>
    );
  }
  return (
      <div>
        <h1>Authorized</h1>
        <Drive/>
        <CreateDrive/>
        <DeleteDrive/>
        <Rename/>
        <Download/>
      </div>
  );
}

export default App;
