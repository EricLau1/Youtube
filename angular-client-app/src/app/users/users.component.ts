import { Component, OnInit } from '@angular/core';

// add imports
import { Router } from '@angular/router';
import { UsersService } from '../services/users.service';
import { AlertService } from '../utils/alert.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  private users: Array<any> = [];
  private token = `bearer ${localStorage.getItem('token')}`;

  constructor(private router: Router, private usersService: UsersService,
    private alert: AlertService) { }

  async ngOnInit() {

      try {

          const results = await this.usersService.getUsers(this.token).toPromise();
    
          this.users = results;
          console.log(this.users);
      
      } catch (error) {

        localStorage.removeItem('token');
        this.router.navigate(['/']);

      }
  
  }

  private logout() {

    localStorage.removeItem('token');
    this.router.navigate(['/login']);
    this.alert.info("Você saiu da conta!", 'Log Out!');
  
  }

  private delete(id) {

    this.usersService.delete(id, this.token)
          .subscribe(rs => {
            console.log(rs);
            this.users = this.users.filter(user => id != user.id);
            this.alert.error("Usuário deletado da lista!", 'Excluindo usuário!');

            if(this.users.length < 1) {
              localStorage.removeItem('token');
              this.router.navigate(['/']);
            }
            
          });

  }

}
