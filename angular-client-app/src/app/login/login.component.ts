import { Component, OnInit } from '@angular/core';

// add imports
import { Router } from '@angular/router';
import { FormGroup } from '@angular/forms';
import { UsersService } from '../services/users.service';
import { AlertService } from '../utils/alert.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  private user: any = {};

  constructor(private usersService: UsersService, private router: Router,
    private alert: AlertService) { }

  ngOnInit() {

    this.user = {
      email: '',
      password: ''
    };

  }

  async signin(form: FormGroup) {

    if(form.valid) {

        try {

            const response = await this.usersService.login(this.user).toPromise();  

            if(response['token']) {
      
                const token = response['token'];

                localStorage.setItem('token', token);
        
                this.router.navigate(['/users']);
                //console.log('Seu token eh: ' + token);
                this.alert.success('Usuário autenticado com sucesso!', 'Você está logado!');
            
                return;
            
            }

        } catch (error) {

          console.error(error);

        }

    }

    return this.alert.warning('Dados incorretos!', 'Falha ao tentar logar!');

  }

}
