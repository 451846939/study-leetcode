/*
请你给一个停车场设计一个停车系统。停车场总共有三种不同大小的车位：大，中和小，每种尺寸分别有固定数目的车位。

请你实现ParkingSystem类：

ParkingSystem(int big, int medium, int small)初始化ParkingSystem类，三个参数分别对应每种停车位的数目。
bool addCar(int carType)检查是否有carType对应的停车位。carType有三种类型：大，中，小，分别用数字1，2和3表示。一辆车只能停在carType对应尺寸的停车位中。如果没有空车位，请返回false，否则将该车停入车位并返回true。


示例 1：

输入：
["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
[[1, 1, 0], [1], [2], [3], [1]]
输出：
[null, true, true, false, false]

解释：
ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
parkingSystem.addCar(1); // 返回 true ，因为有 1 个空的大车位
parkingSystem.addCar(2); // 返回 true ，因为有 1 个空的中车位
parkingSystem.addCar(3); // 返回 false ，因为没有空的小车位
parkingSystem.addCar(1); // 返回 false ，因为没有空的大车位，唯一一个大车位已经被占据了


提示：

0 <= big, medium, small <= 1000
carType取值为1，2或3
最多会调用addCar函数1000次

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-parking-system
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
struct ParkingSystem {
    park: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl ParkingSystem {

    fn new(big: i32, medium: i32, small: i32) -> Self {
        ParkingSystem { park: vec![0, big, medium, small] }
    }

    fn add_car(&mut self, car_type: i32) -> bool {
        self.park[car_type as usize] -= 1;
        self.park[car_type as usize] >= 0
    }
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * let obj = ParkingSystem::new(big, medium, small);
 * let ret_1: bool = obj.add_car(carType);
 */