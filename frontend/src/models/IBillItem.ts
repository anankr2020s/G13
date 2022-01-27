import { ExaminationInterface } from "./IExamination";
import { BillInterface } from "./IBill";
export interface BillitemInterface{
    ID : number,
    ExaminationID:number,
    Examination : ExaminationInterface,
    BillID : number,
    Bill : BillInterface
}