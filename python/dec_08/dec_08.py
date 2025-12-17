import math

class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.size = [1] * n
    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]
    def union(self, x, y):
        xr, yr = self.find(x), self.find(y)
        if xr == yr:
            return
        if self.size[xr] < self.size[yr]:
            self.parent[xr] = yr
            self.size[yr] += self.size[xr]
        else:
            self.parent[yr] = xr
            self.size[xr] += self.size[yr]
    def sizes(self):
        roots = [self.find(x) for x in range(len(self.parent))]
        counts = {}
        for r in roots:
            counts[r] = counts.get(r, 0) + 1
        return list(counts.values())

def first():
    with open("input_08.txt") as f:
        boxes = [tuple(map(int, line.strip().split(","))) for line in f]
    n = len(boxes)
    pairs = []
    for i in range(n):
        for j in range(i+1, n):
            a, b = boxes[i], boxes[j]
            dist = math.sqrt((a[0]-b[0])**2 + (a[1]-b[1])**2 + (a[2]-b[2])**2)
            pairs.append((dist, i, j))
    pairs.sort()
    uf = UnionFind(n)
    connections = 0
    for idx in range(1000):
        _, i, j = pairs[idx]
        uf.union(i, j)
    sizes = sorted(uf.sizes(), reverse=True)
    while len(sizes) < 3:
        sizes.append(1)
    result = sizes[0] * sizes[1] * sizes[2]
    print(f"first: {result} (sizes: {sizes[0]}, {sizes[1]}, {sizes[2]})")

def second():
    with open("input_08.txt") as f:
        boxes = [tuple(map(int, line.strip().split(","))) for line in f]
    n = len(boxes)
    pairs = []
    for i in range(n):
        for j in range(i+1, n):
            a, b = boxes[i], boxes[j]
            dist = math.sqrt((a[0]-b[0])**2 + (a[1]-b[1])**2 + (a[2]-b[2])**2)
            pairs.append((dist, i, j))
    pairs.sort()
    uf = UnionFind(n)
    last_i = last_j = None
    for _, i, j in pairs:
        if uf.find(i) != uf.find(j):
            uf.union(i, j)
            last_i, last_j = i, j
            if len(set(uf.find(x) for x in range(n))) == 1:
                break
    if last_i is not None and last_j is not None:
        result = boxes[last_i][0] * boxes[last_j][0]
        print(f"second: {result} (X: {boxes[last_i][0]}, {boxes[last_j][0]})")
    else:
        print("second: not found")

if __name__ == "__main__":
    first()
    second()
