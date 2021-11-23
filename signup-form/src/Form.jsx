import React from "react";
import './Form.css';
class Signup extends React.Component {
    constructor(props) {
      super(props);
      this.state = {};
      this.handleInputChange = this.handleInputChange.bind(this);
    }
  
    handleInputChange(event) {
      const target = event.target;
      const name = target.name === 'name'? target.value: ""
      const domain = target.name === 'domain'? target.value: ""
  
      this.setState({
        name : name,
        domain: domain
      });
    }
    
    render() {
      return (
        <form className= "form-style">
        <label>
            Name:
        <input type="text" onChange={this.handleInputChange} name="name"/>
        </label>  
        <br/>  
        <label>
            Domain:
        <input type="text" onChange={this.handleInputChange} name="domain"/>
        </label>
        <br/>
        <input type="submit" value="Submit" />
        </form>
      );
    }
  }

export default Signup;