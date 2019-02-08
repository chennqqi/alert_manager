


const url_alerts = 'api/alerts';
const url_auth = 'api/auth';

// var request = require('request');

function handleErrors(response) {
    if (!response.ok) {
        throw Error(response.statusText);
    }
    return response;
}

export class AlertManagerApi {

    constructor(url) {
        // this.name = name;
        this.url = url;
        this.token = null;
        // this.getWaitApiToken()
        this.getApiToken();
    }

    // Adding a method to the constructor
    getAlertsList({
        limit=500, 
        aggregate=true, 
        timerange_h=96, 
        sites=[], 
        devices=[],
        status=[1,2,3]}={}) {

        var params = `?limit=${limit}`

        if (aggregate) {
            params = params + `&agg_id=0`
        }

        // if (timerange_h) {
        //     params = params + `&timerange=72h`
        // }

        if (sites.length !== 0) {
            params = params + `&site__in=${sites.join(',')}`
        }

        if (devices.length !== 0) {
            params = params + `&device__in=${devices.join(',')}`
        }

        if (status.length !== 0) {
            params = params + `&status__in=${status.join(',')}`
        }

        console.log("fetching > " + this.url + url_alerts + params)
        return fetch(this.url + url_alerts + params)
          .then(response => response.json());
    }

    getAlert(id) {
        return fetch(`${this.url}${url_alerts}/${id}`)
          .then(response => response.json());
    }
    getAlertWithHistory(id) {
        return fetch(`${this.url}${url_alerts}?id=${id}&history=true`)
          .then(response => response.json())
          .then(data => data[0]);
    }

    bulkUpdateStatus({items, status}={}) {

        for( var i in items) {
           console.log(`Will update status for ${i} with ${status}`)
       }
   }

    getContributingAlerts(id) {
        return fetch(`${this.url}${url_alerts}?agg_id=${id}` )
          .then(response => response.json());
    }

    updateAlertOwner({id, owner, team}={}) {

        let url = `${this.url}${url_alerts}/${id}?owner=${owner}&team=${team}`
        
        this.getWaitApiToken()

        console.log("Token is : " + this.token)
        let obj = {
            method: 'PATCH',
            headers: {
                // "Content-Type": "application/json",
                "Authorization": `Bearer ${this.token}`
            }
        }

        return fetch( url, obj )
          .then(handleErrors)
          .catch(function(error) {
            console.log(error);
        });
    }

    updateAlertStatus({id, status}={}) {

        let status_to_id = {
            ACTIVE:  1,
            SUPPRESSED: 2,
            EXPIRED: 3,
            CLEARED: 4
        }
        
        if (!(status.toUpperCase() in status_to_id)) {
            console.log(`Unable to update status for ${id}, Status: ${status} is not supported ..`);
            return
        }
        
        let url = `${this.url}${url_alerts}/${id}?status=${status_to_id[status]}`

        let obj = {
            method: 'PATCH',
            headers: {
                // "Content-Type": "application/json",
                "Authorization": `Bearer ${this.token}`
            }
        }

        return fetch( url, obj )
          .then(handleErrors)
          .catch(function(error) {
            console.log(error);
        });
    }

    updateAlertSeverity({id, severity}={}) {

        let severity_to_id = {
            CRITICAL:  1,
            WARN: 2,
            INFO: 3,
        }

        if (!(severity.toUpperCase() in severity_to_id)) {
            console.log(`Unable to update severity for ${id}, Severity: ${severity} is not supported ..`);
            return
        }
        
        let url = `${this.url}${url_alerts}/${id}?severity=${severity_to_id[severity]}`
        
        let obj = {
            method: 'PATCH',
            headers: {
                // "Content-Type": "application/json",
                "Authorization": `Bearer ${this.token}`
            }
        }

        return fetch( url, obj )
          .then(handleErrors)
          .catch(function(error) {
            console.log(error);
        });
    }

    alertClear({id}={}) {
        // api/alerts/{id}/clear
    
        let url = `${this.url}${url_alerts}/${id}/clear`
        
        let obj = {
            method: 'PATCH',
            headers: {
                // "Content-Type": "application/json",
                "Authorization": `Bearer ${this.token}`
            }
        }

        return fetch( url, obj )
          .then(handleErrors)
          .catch(function(error) {
            console.log(error);
        });

    }

    alertSuppress({id, duration="1h"}={}) {

        // api/alerts/{id}/suppress?duration=5m
        let url = `${this.url}${url_alerts}/${id}/suppress?duration=${duration}`
        
        let obj = {
            method: 'PATCH',
            headers: {
                // "Content-Type": "application/json",
                "Authorization": `Bearer ${this.token}`
            }
        }

        return fetch( url, obj )
          .then(handleErrors)
          .catch(function(error) {
            console.log(error);
        });
    }

    alertAcknowledge({id, owner="owner", team="team"}={}) {
        // api/alerts/{id}/acknowledge?owner=foo&team=bar
        let url = `${this.url}${url_alerts}/${id}/acknowledge?owner=${owner}&team=${team}`
        
        let obj = {
            method: 'PATCH',
            headers: {
                // "Content-Type": "application/json",
                "Authorization": `Bearer ${this.token}`
            }
        }

        return fetch( url, obj )
          .then(handleErrors)
          .catch(function(error) {
            console.log(error);
        });
    }

    async getWaitApiToken() {

        if (this.token === null) {
            console.log("WIll query a new token");
            try {
                var response = await fetch(this.url + url_auth, {
                                                method: 'POST',
                                                body: JSON.stringify({ username: "react", password: "react" })
                                            })
                var resp = await response.json();
                this.token = resp.token
                console.log("Token Fetched: " + this.token)
            } catch (err) {
                console.log('fetch failed', err);
            }
        } else {
            console.log("Token already present");
        }

              
        //     var resp = await Promise((resolve, reject) => {
        //         var token = fetch(this.url + url_auth, {
        //             method: 'POST',
        //             body: JSON.stringify({ username: "react", password: "react" })
        //           })
        //           .then(handleErrors)
        //           .then(response => response.json())
            
        //         resolve(token)
    
        //     });

        //     this.token = resp.token
        //     console.log("Queried token: " + this.token);
        // } else {
        //     console.log("Token already present");
        // }
    }

    getApiToken() {

    //     return new Promise((resolve, reject) => {
    //         var token = fetch(this.url + url_auth, {
    //             method: 'POST',
    //             body: JSON.stringify({ username: "react", password: "react" })
    //           })
    //           .then(handleErrors)
    //           .then(response => response.json());
        
    //         resolve(token);

    //       });

        return fetch(this.url + url_auth, {
            method: 'POST',
            body: JSON.stringify({ username: "react", password: "react" })
          })
          .then(handleErrors)
          .then(response => response.json())
          .then(data => this.setApiToken( data ) )
          .catch(function(error) {
            console.log(error);
            });
        
    }

    setApiToken( data ) {
        this.token = data.token;
        console.log(this.token);
    }
}

