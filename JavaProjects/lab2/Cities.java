public class Cities {
    private int n;
    private int [][] roads;
    public Cities(int k){
        roads = new int[k][k];
        for (int i = 0; i < k; i++){
            for (int j = 0; j < k; j++){
                if (i == j)
                    this.roads[i][j] = 0;
                else
                    this.roads[i][j] = -1;
            }
        }
        n = k;
    }
    public void setRoad(int a, int b, int l){
        System.out.println("дорога из *" + a + "* в *" +
                b +"* равна: " + l);
        this.roads[a][b] = l;
        this.roads[b][a] = l;
    }
    public void outMatrix(){
        for (int i = 0; i < n; i++){
            for (int j = 0; j < n; j++) {
                System.out.print(roads[i][j] + " ");
            }
            System.out.println();
        }
    }
    public int path(int[] nums){
        int sum = 0;
        System.out.print("пройдём по пути городов под номерами: ");
        for (int i = 0; i < nums.length - 1; i++){
            System.out.print(nums[i] + "->");
        }
        System.out.println(nums[nums.length - 1]);
        for (int i = 0; i < nums.length - 1; i++){
            if (this.roads[nums[i]][nums[i + 1]] < 0) {
                System.out.print("дороги из города *"+nums[i]);
                System.out.println("* в город *"+nums[i + 1]+"* не существует,");
                System.out.print("пожалуйста, создайте дорогу или ");
                System.out.println("выберите другую последовательность городов.");
                return -1;
            }
            sum += this.roads[nums[i]][nums[i + 1]];
        }
        return sum;
    }
    public int getRoad(int a, int b){
        return roads[a][b];
    }
}
