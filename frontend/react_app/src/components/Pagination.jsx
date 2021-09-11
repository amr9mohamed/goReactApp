import React from "react";

export const Pagination = ({ elementsPerPage, totalElements, paginate }) => {
  const pageNumbers = [];
  for (let i = 1; i <= Math.ceil(totalElements / elementsPerPage); i++) {
    pageNumbers.push(i);
  }
  return (
    <nav>
      <ul className="pagination">
        {pageNumbers.map((number) => (
          <li key={number} className="page-item">
            <a href="!#" className="page-link" onClick={() => {paginate(number)}}>
              {number}
            </a>
          </li>
        ))}
      </ul>
    </nav>
  );
};
