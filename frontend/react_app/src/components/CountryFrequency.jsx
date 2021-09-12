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

export const CountryFrequency = ({ stats }) => {
  return (
    <div style={{ height: 370, width: "40%" }}>
      <DataGrid
        rows={stats}
        columns={columns}
        pageSize={5}
        disableSelectionOnClick
      />
    </div>
  );
};
