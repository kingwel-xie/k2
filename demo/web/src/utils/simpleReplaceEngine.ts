interface Operator {
  op(val: string | undefined): string;
}
class Sub implements Operator {
  offset: number;

  length?: number;

  constructor(token: string) {
    const hasLen = token.indexOf(':');
    if (hasLen == -1) {
      this.offset = parseInt(token);
    } else {
      const [offset, length] = token.split(':', 2);
      this.offset = parseInt(offset);
      this.length = parseInt(length);
    }
  }
  op(val: string | undefined): string {
    if (!val) return '';
    return val.slice(this.offset, this.length);
  }
}

interface Builder {
  build(data: any): string;
}

class Text implements Builder {
  constructor(private text: string) {}
  build(_data: any): string {
    return this.text;
  }
}

class Replacer implements Builder {
  constructor(private field: string, private op?: Operator) {}
  build(data: any): string {
    const val = data[this.field];
    return this.op ? this.op.op(val) : val;
  }
}

export function replace(template: string, data: any): string {
  const builders: Builder[] = [];
  do {
    const start = template.indexOf('{');
    if (start === -1) {
      break;
    }
    const end = template.indexOf('}', start + 1);
    if (end === -1) {
      break;
    }
    if (start !== 0) {
      builders.push(new Text(template.slice(0, start)));
    }
    const replace = template.slice(start + 1, end);
    const [field, op] = parseFiledAndOp(replace);
    builders.push(new Replacer(field, op));
    template = template.slice(end + 1);
  } while (template.length);
  if (template.length) {
    builders.push(new Text(template));
  }
  return builders.map((builder) => builder.build(data)).join('');
}

function parseFiledAndOp(replace: string): [string, Operator?] {
  let op = -1;
  if ((op = replace.indexOf(':')) !== -1) {
    return [replace.slice(0, op), new Sub(replace.slice(op + 1))];
  }
  return [replace];
}
