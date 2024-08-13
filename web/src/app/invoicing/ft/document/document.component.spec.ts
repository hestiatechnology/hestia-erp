import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FaturaDocumentComponent } from './document.component';

describe('DocumentComponent', () => {
  let component: FaturaDocumentComponent;
  let fixture: ComponentFixture<FaturaDocumentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FaturaDocumentComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(FaturaDocumentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
