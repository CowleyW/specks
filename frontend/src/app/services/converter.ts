export interface Converter {
  convert(jsonData: any[]): string;
}

export enum Format {
  CSV = "CSV",
  JSON = "JSON",
  SQL = "SQL"
}

export class ConverterFactory {
  public static createConverter(format: Format): Converter | null {
    switch (format) {
      case Format.CSV:
        return new DelimitedCSVConverter();
      case Format.JSON:
        return new JSONConverter();
      default:
        console.log("Can't convert this type");
        return null;
    }
  }
}

class DelimitedCSVConverter implements Converter {
  public convert(jsonData: any[]): string {
    if (jsonData.length === 0) {
      return '';
    }

    return '';
  }
}

class JSONConverter implements Converter {
  public convert(jsonData: any[]): string {
    return JSON.stringify(jsonData);
  }
}
