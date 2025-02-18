
export interface AdditionalData {
    algorithm: string; // Ej.: "AES"
    mode: string;      // Ej.: "GCM"
    strength: number;  // Ej.: 256
    iv64: string;      // Ej.: "32bVr0KW+Cj5pPLB"
  }



export interface Payload {
  cypher64: string;
  additionaldata: AdditionalData;
}



export interface Err {
    Error(): string;
}


export type MayErr = Err | null;



export type MayAdditionalData = Buffer | null;

export interface Result {
    Data: Buffer;
    MayErr: MayErr;

}


export function SimpleErr(err: string): Err {
    return {
        Error: () => err
    };
}