import React from 'react'
import '../styles/Header.css'
import {generalStrings} from "../strings/localized-strings"

class Header extends React.Component {
  render() {
    return (
      <div className="headerContainer">
        <h1 className="defaultTitle">{generalStrings.siteTitle}</h1>
      </div>
    )
  }
}

export default Header
