import { Component, OnInit } from '@angular/core';

// add imports
import { ActivatedRoute } from '@angular/router';
import { FormGroup } from '@angular/forms';
import { UsersService } from '../../services/users.service';
import { AlertService } from '../../utils/alert.service';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {

  private user: any = {};
  private token = `bearer ${localStorage.getItem('token')}`;

  constructor(private router: ActivatedRoute, private usersService: UsersService,
    private alert: AlertService) { }

  async ngOnInit() {

      this.user = {
        id: null,
        name: '',
        email: '',
        password: ''
      };

      this.router.params.subscribe(param => {
        
        let id = param.id;
        console.log('seu id eh: ' + id);

        this.usersService.getUser(id, this.token)
          .subscribe(data => {
            this.user = data;
            console.log(this.user);
          });

      });
    
  }

  async onSave(form: FormGroup) {

    if(form.valid) {
      
      try {

        const response = await this.usersService.update(this.user, this.token).toPromise();

        if(response == 1) {
          console.log("OK!");
          this.alert.info('As informações foram atualizadas com sucesso!', 'Aviso');
          return;
        }
      
      } catch (error) {

        console.error(error);

      }
  
    } 

    return this.alert.error('Os dados não foram atualizados!', 'Falha na atualização.');

  }

}
