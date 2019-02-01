export type ErrorMeta = {
  [key: string]: string
};

export type Error = {
  code: string
  message: string
  meta?: ErrorMeta
};

export interface ValidationErrors {
  [key: string]: string;
}
