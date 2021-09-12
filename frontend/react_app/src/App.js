import React, { useState, useEffect } from "react";
import axios from "axios";
import { Users } from "./components/Users";
import { CountryFrequency } from "./components/CountryFrequency";
import { makeStyles } from "@material-ui/core/styles";
import InputLabel from "@material-ui/core/InputLabel";
import FormControl from "@material-ui/core/FormControl";
import Select from "@material-ui/core/Select";
import Button from "@material-ui/core/Button";
import "./App.css";

const useStyles = makeStyles((theme) => ({
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
}));

function App() {
  const classes = useStyles();
  const [users, setUsers] = useState([]);
  const [loading, setloading] = useState(false);
  const [usersPerPage] = useState(10);
  const [countries, setCountries] = useState([]);
  const [selectedCountry, setSelectedCountry] = useState("");
  const [states, setStates] = useState([]);

  useEffect(() => {
    const fetchCountries = async () => {
      setloading(true);
      const res = await axios.get("http://localhost:8080/users/countries");
      setCountries(res.data);
      setloading(false);
    };
    fetchCountries();
  }, []);

  const handleChange = (event) => {
    setSelectedCountry(event.target.value);
  };

  const handleGetStates = () => {
    axios
      .get("http://localhost:8080/users/frequency")
      .then((res) => {
        res.data = res.data.map((el, idx) => {
          return { ...el, id: idx };
        });
        setStates(res.data);
      })
      .catch((err) => console.log(err));
  };

  return (
    <div>
      <h1 className="text-primary mb-3">Users</h1>
      <FormControl className={classes.formControl}>
        <InputLabel htmlFor="country-native-simple">Country</InputLabel>
        <Select
          native
          value={selectedCountry}
          onChange={handleChange}
          inputProps={{
            name: "country",
            id: "country-native-simple",
          }}
        >
          <option aria-label="None" value="" />
          {countries.map((country) => (
            <option key={country} value={country}>
              {country}
            </option>
          ))}
        </Select>
      </FormControl>
      <br />
      {users.length ? (
        <Users users={users} loading={loading} usersPerPage={usersPerPage} />
      ) : null}
      <br />
      <Button variant="contained" color="primary" onClick={handleGetStates}>
        Show Stats
      </Button>
      {states.length ? (
        <CountryFrequency states={states} loading={loading} />
      ) : null}
    </div>
  );
}

export default App;
