const STACK_CAPACITY = 100;

export function push<T>(stack: T[], item: T): T[] {
    if (stack.length === STACK_CAPACITY) {
        stack.shift();
    }
    stack.push(item);
    return stack;
}

export function removeLast<T>(stack: T[]): T[] {
    if (stack.length > 1) {
        stack.pop();
    }
    return stack;
}

export function peek<T>(stack: T[]): T | undefined {
    return stack[stack.length - 1];
}

export function peek2<T>(stack: T[]): T | undefined {
    return stack[stack.length - 2];
}
