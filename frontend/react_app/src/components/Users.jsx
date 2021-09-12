import React from "react";
import { DataGrid } from "@material-ui/data-grid";

const columns = [
  { field: "id", headerName: "ID", width: 120 },
  {
    field: "email",
    headerName: "Email",
    width: 200,
  },
  {
    field: "country",
    headerName: "Country",
    width: 200,
  },
  {
    field: "phoneNumber",
    headerName: "Phone Number",
    width: 200,
  },
  {
    field: "parcelWeight",
    headerName: "Parcel Weight",
    width: 175,
  },
];

export const Users = ({ users, usersPerPage }) => {
  return (
    <div style={{ height: 650, width: "100%" }}>
      <DataGrid
        rows={users}
        columns={columns}
        pageSize={usersPerPage}
        disableSelectionOnClick
      />
    </div>
  );
};
