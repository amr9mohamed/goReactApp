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
  const [usersPerPage, SetUsersPerPage] = useState(10);
  const [countries, setCountries] = useState([]);
  const [selectedCountry, setSelectedCountry] = useState("");
  const [stats, setStats] = useState([]);
  const [numberOfUsers, setNumberOfUsers] = useState(0);
  const [currentPage, setCurrentPage] = useState(1);
  const [showStatsFlag, setShowStatsFlag] = useState(false);

  useEffect(() => {
    const fetchCountries = async () => {
      const res = await axios.get("http://localhost:8080/users/countries");
      setCountries(res.data);
    };
    fetchCountries();
  }, []);

  useEffect(() => {
    const fetchFrequencies = async () => {
      const res = await axios.get("http://localhost:8080/users/frequency");
      res.data = res.data.map((el, idx) => {
        return { ...el, id: idx };
      });
      setStats(res.data);
    };
    fetchFrequencies();
  }, []);

  const handleChange = (event) => {
    let targetCountry = event.target.value;
    if (targetCountry) {
      setSelectedCountry(targetCountry);
      setCurrentPage(1);
      setloading(true);
      setNumberOfUsers(
        stats.find((el) => el.country === targetCountry).frequency
      );
      axios
        .get(
          `http://localhost:8080/users/${targetCountry}/${currentPage}/${usersPerPage}/`
        )
        .then((res) => {
          setUsers(res.data);
        })
        .catch((err) => console.log(err))
        .finally(() => {
          setloading(false);
        });
    }
  };

  const handlePageChange = (event, newPage) => {
    setCurrentPage(newPage + 1);
    axios
      .get(
        `http://localhost:8080/users/${selectedCountry}/${
          newPage + 1
        }/${usersPerPage}/`
      )
      .then((res) => {
        setUsers(res.data);
      })
      .catch((err) => console.log(err))
      .finally(() => {
        setloading(false);
      });
  };

  const handleChangeRowsPerPage = (event) => {
    SetUsersPerPage(+event.target.value);
    axios
      .get(
        `http://localhost:8080/users/${selectedCountry}/${currentPage}/${+event
          .target.value}/`
      )
      .then((res) => {
        setUsers(res.data);
      })
      .catch((err) => console.log(err))
      .finally(() => {
        setloading(false);
      });
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
        <Users
          users={users}
          page={currentPage - 1}
          numberOfUsers={numberOfUsers}
          usersPerPage={usersPerPage}
          handlePageChange={handlePageChange}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      ) : (
        [loading ? <h2>Loading...</h2> : null]
      )}
      <br />
      <Button
        variant="contained"
        color="primary"
        onClick={() => setShowStatsFlag(true)}
      >
        Show Stats
      </Button>
      {showStatsFlag ? <CountryFrequency stats={stats} /> : null}
    </div>
  );
}

export default App;
