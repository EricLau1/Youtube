import { Injectable } from '@angular/core';

// add imports
import { HttpClient, HttpHeaders } from '@angular/common/http';

const ApiRoutes = {
  login: 'login',
  signup: 'signup',
  users: 'users'
};

@Injectable({
  providedIn: 'root'
})
export class UsersService {

  private UriApi: string = 'http://localhost:3000';

  constructor(private http: HttpClient) { }

  private loadHeaders(token: string = '') {

    let headers = new HttpHeaders({
      'Content-type': 'application/json',
      'Authorization': `${token}`
    });

    return { headers };

  }

  public login(user: any) {

    let uri = `${this.UriApi}/${ApiRoutes.login}`;

    return this.http.post(uri, JSON.stringify(user), this.loadHeaders() );

  }

  public create(user: any) {

    let uri = `${this.UriApi}/${ApiRoutes.signup}`;

    return this.http.post(uri, JSON.stringify(user), this.loadHeaders() );

  }


  public getUsers(token: string = '') {

    let uri = `${this.UriApi}/${ApiRoutes.users}`;

    return this.http.get<Array<any>>(uri, this.loadHeaders(token));

  }

  public getUser(id, token: string = '' ) {

    let uri = `${this.UriApi}/${ApiRoutes.users}/${id}`;

    return this.http.get<any>(uri, this.loadHeaders(token));

  }

  public update(user: any, token: string = '') {

    let uri = `${this.UriApi}/${ApiRoutes.users}/${user.id}`;

    return this.http.put(uri, JSON.stringify(user), this.loadHeaders(token));

  }


  public delete(id, token: string = '') {

    let uri: string = `${this.UriApi}/${ApiRoutes.users}/${id}`;
    
    return this.http.delete(uri, this.loadHeaders(token));

  }



}
