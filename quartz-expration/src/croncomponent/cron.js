import React, { Component } from 'react'
import { ReCron } from '@sbzen/re-cron';
import axios from 'axios';

class Cron extends Component {
    constructor(props) {
        super(props);
        this.state = {
            cronValue: '0 0 12 L-2 * ? *',
            description:'At 12:00, 2 days before the last day of the month',
            nextFiveSchedule:["Fri Jan 29 12:00:00 2021","Fri Feb 26 12:00:00 2021","Mon Mar 29 12:00:00 2021","Wed Apr 28 12:00:00 2021","Sat May 29 12:00:00 2021"]
        };
    }
    handleChange(cronVal) {
        this.setState({ cronValue:cronVal });
    }
    getDescription() {
        axios.get('/api/description?expration='+this.state.cronValue)
        .then(response=>{
            this.setState({ description:response.data.data });
        })
        .catch(error=>{
            console.error(error)
        })
        
    }
    getNextFiveScheduleTime() {
        axios.get('/api/next/5?expration='+this.state.cronValue)
        .then(response=>{
            this.setState({ nextFiveSchedule:response.data.data });
        })
        .catch(error=>{
            console.error(error)
        })
        
    }
    render() {
        const {nextFiveSchedule} = this.state
        return (
            
            <div className="container-fluid p-0 m-0">
                <div className="jumbotron p-3 m-5">
                  <h3 className="display-6">Quartz Cron Expration</h3>
                  <hr className="my-4"></hr>
                  <div className="row bg-light text-dark p-5 m-1 rounded-lg">
                 
                      <div className="col col-6">
                            <ReCron 
                                value={this.state.cronValue}
                                onChange={(e) => this.handleChange(e)}>
                            </ReCron>
                      </div>
                      <div className="col col-6">
                        <button type="button" className="btn btn-secondary" onClick={()=>this.getDescription()}>Describe</button> &nbsp;
                        <button type="button" className="btn btn-secondary" onClick={()=>this.getNextFiveScheduleTime()}>NextFiveSchedule</button> 
                        <span className="display-5 font-weight-bold pl-3">Expration : {this.state.cronValue}</span>
                        <hr className="my-1"></hr>
                        <div className="p-2 m-3"> 
                            <p className="lead"><strong>Description : </strong>{this.state.description}.</p>
                            <div>
                                {
                                    nextFiveSchedule.length?
                                    nextFiveSchedule.map((schedule) =><p className="text-monospace text-uppercase" key={schedule}>{schedule}</p>) :
                                    null
                                }
                            </div>
                         </div>   
                      </div>
                  </div>
                                  
                </div>
            </div>
        );
    }
}

export default Cron
