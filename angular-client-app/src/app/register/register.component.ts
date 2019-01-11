import { Component, OnInit } from '@angular/core';

import { FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { UsersService } from '../services/users.service';
import { AlertService } from '../utils/alert.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  private user: any = {};

  constructor(private usersService: UsersService, private router: Router,
    private alert: AlertService) { }

  ngOnInit() {

    this.user = {
      name: '',
      email: '',
      password: ''
    };

  }

  async signup(form: FormGroup) {

    if(form.valid) {
        
      try {

        const response = await this.usersService.create(this.user).toPromise();

        localStorage.removeItem('token');

        this.router.navigate(['/login']);

        this.alert.success('Você foi cadastrado com sucesso!', 'Cadastro salvo!');

        return;

      } catch(error) {

        console.error(error);

      }

    } // end if

    return this.alert.warning('Os dados estão incorretos!', 'Erro ao cadastrar!');

  }

}
