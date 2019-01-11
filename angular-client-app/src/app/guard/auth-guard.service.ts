import { Injectable } from '@angular/core';

// add imports
import { CanActivate, Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate{

  constructor(private router: Router) { }

  canActivate() {

    const token = localStorage.getItem('token') || null;

    if(token != null) {

      return true;

    }

    this.router.navigate(['/login']);

    return false;

  }

}
