import { Injectable } from '@angular/core';

// add imports
import { ToastrManager } from 'ng6-toastr-notifications';

@Injectable({
  providedIn: 'root'
})
export class AlertService {

  constructor(public toastr: ToastrManager) { }

  public success(message, title: string = 'Success!') {
    this.toastr.successToastr(message, title);
  }

  public error(message, title: string = 'Error!') {
    this.toastr.errorToastr(message, title);
  }

  public warning(message, title: string = 'Warning!') {
    this.toastr.warningToastr(message, title);
  }

  public info(message, title: string = 'Alert!') {
    this.toastr.infoToastr(message, title);
  }

}
