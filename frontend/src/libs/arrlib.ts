function isSimilar<T>(a1: T[], a2: T[]): boolean {
  if (! Array.isArray(a1) || ! Array.isArray(a2)) {
    return false
  }
  if (a1.length !== a2.length) {
    return false
  }

  const set1 = new Set(a1)
  const set2 = new Set(a2)

  return a1.every(item => set2.has(item)) && a2.every(item => set1.has(item))
}

export default {
  isSimilar,
}
