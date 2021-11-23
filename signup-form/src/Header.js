import React from 'react';
import './Header.css';
import logoUrl from './assets/logo.png';

// const handleFilterChange = e => {
//   const value = e.target.value || undefined;
//   setFilter("OBJ_EXEC_STATUS", value);
//   setFilterInput(value);
// };

const Header = () => (
  <div className="header">
    <img src={logoUrl} alt="Filing Box" />
    <input
        // value={filterInput}
        // onChange={handleFilterChange}
        placeholder={"Search for Products and features"}
      />
  </div>
)
export default Header;