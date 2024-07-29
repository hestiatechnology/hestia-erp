import { Component, Input, OnInit } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { RouterModule } from "@angular/router";
import { Link, NavbarService } from "./navbar.service";
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatTreeNestedDataSource, MatTreeModule } from '@angular/material/tree';
import { NestedTreeControl } from "@angular/cdk/tree";



@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [MatToolbarModule, MatButtonModule, MatIconModule, MatSidenavModule, MatListModule, RouterModule, MatFormFieldModule, MatTreeModule],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss'
})

export class NavbarComponent implements OnInit {
  @Input() title: string = "Hestia ERP";
  @Input() links: Link[] = [];


  constructor(private navbarService: NavbarService) {
    this.dataSource.data = this.links;
  }

  ngOnInit() {
    this.navbarService.title$.subscribe(title => {
      this.title = title;
    });

    /* this.navbarService.links$.subscribe(links => {
      this.links = links;
    }); */
  }

  treeControl = new NestedTreeControl<Link>(node => node.children);
  dataSource = new MatTreeNestedDataSource<Link>();


  hasChild = (_: number, node: Link) => !!node.children && node.children.length > 0;
}
