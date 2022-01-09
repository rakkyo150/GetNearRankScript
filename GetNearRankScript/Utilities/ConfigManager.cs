using Newtonsoft.Json;

internal class ConfigManager
{
    internal void LoadConfigFile(string path)
    {
        StreamReader re = new StreamReader(path);
        string _jsonStr=re.ReadToEnd();
        re.Close();
        dynamic _jsonDyn=JsonConvert.DeserializeObject<dynamic>(_jsonStr);

        if (_jsonDyn==null||_jsonDyn.YourId=="")
        {
            Console.WriteLine("There is wrong with Config.json");
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

        string _jsonFinish=JsonConvert.SerializeObject(Config.Instance, Formatting.Indented);
        
        StreamWriter wr = new StreamWriter(path, false);
        wr.WriteLine(_jsonFinish);
        wr.Close();
    }
}
