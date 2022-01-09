using Newtonsoft.Json;

internal class ConfigManager
{
    internal void LoadConfigFile(string path)
    {
        StreamReader re = new StreamReader(path);
        string _jsonStr=re.ReadToEnd();
        re.Close();
        dynamic _jsonDyn=JsonConvert.DeserializeObject<dynamic>(_jsonStr);

        if (_jsonDyn==null)
        {
            Console.WriteLine("There is something wrong with Config.json");
            Console.WriteLine("Remake Config.json");
            MakeConfigFile(path);
            LoadConfigFile(path);
            return;
        }
        
        if (_jsonDyn != null)
        {
            Config.Instance.YourId = _jsonDyn.YourId;
            Config.Instance.RankRange = _jsonDyn.RankRange;
            Config.Instance.PPFilter= _jsonDyn.PPFilter;
            Config.Instance.YourPageRange = _jsonDyn.YourPageRange;
            Config.Instance.OthersPageRange = _jsonDyn.OthersPageRange;
        }
    }
    
    internal void MakeConfigFile(string path)
    {
        Console.WriteLine("Input Your ID");
        Config.Instance.YourId = Console.ReadLine();
        Console.WriteLine("Input Rank Range");
        Config.Instance.RankRange = int.Parse(Console.ReadLine());
        Console.WriteLine("Input PP Filter");
        Config.Instance.PPFilter = int.Parse(Console.ReadLine());
        Console.WriteLine("Input Your Page Range");
        Config.Instance.YourPageRange = int.Parse(Console.ReadLine());
        Console.WriteLine("Input Rivals' Page Range");
        Config.Instance.OthersPageRange = int.Parse(Console.ReadLine());

        string _jsonFinish=JsonConvert.SerializeObject(Config.Instance, Formatting.Indented);

        // 変えなくてもよかった
        StreamWriter wr = new StreamWriter(new FileStream(path,FileMode.Create));
        wr.WriteLine(_jsonFinish);
        wr.Close();
    }
}
