import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FaturaHomeComponent } from './home.component';

describe('HomeComponent', () => {
  let component: FaturaHomeComponent;
  let fixture: ComponentFixture<FaturaHomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FaturaHomeComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(FaturaHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
