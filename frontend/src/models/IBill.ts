import { CashierInterface } from "./ICashier";
import { PatientInterface } from "./IPatient";
import { PatientRightInterface } from "./IPatientRight";
import { PaytypeInterface } from "./IPaytype";
import { BillitemInterface } from "./IBillItem";

export interface BillInterface{
    ID: number,

    PatientRightID: number,
    PatientRight:   PatientRightInterface,

    BillTime:   Date,
    Total:  number,
    Telephone : string,

    CashierID: number,
    Cashier: CashierInterface,

    PaytypeID : number,
    Paytype : PaytypeInterface,
    
    BillItems : BillitemInterface[]

}