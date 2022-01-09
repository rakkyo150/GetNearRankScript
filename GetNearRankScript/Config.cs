internal class Config
{
    public static Config Instance { get; set; } = new Config();
    public string YourId { get; set; } = "";
    public int RankRange { get; set; } = 3;
    public int PPFilter { get; set; } = 20;
    public int YourPageRange { get; set; } = 10;
    public int OthersPageRange { get; set; } = 3;
}
