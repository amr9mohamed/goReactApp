import React from "react";
import { DataGrid } from "@material-ui/data-grid";

const columns = [
  {
    field: "country",
    headerName: "Country",
    width: 200,
  },
  {
    field: "frequency",
    headerName: "Frequency",
    width: 175,
  },
];

export const CountryFrequency = ({ states, loading }) => {
  if (loading) {
    return <h2>Loading...</h2>;
  }

  return (
    <div style={{ height: 350, width: "40%" }}>
      <DataGrid
        rows={states}
        columns={columns}
        pageSize={5}
        disableSelectionOnClick
      />
    </div>
  );
};
