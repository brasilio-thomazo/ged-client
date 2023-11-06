<?php

namespace App\Console\Commands;

use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\DB;
use Illuminate\Console\Command;
use App\Models\DocumentType;
use App\Models\Department;
use App\Models\Search;
use App\Models\Group;
use App\Models\User;
use Exception;

class Install extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'app:install';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'install client database';

    /**
     * Execute the console command.
     */
    public function handle()
    {
        $buffer = "";

        $rw = ['group' => 'rw', 'user' => 'rw', 'search' => 'rw', 'document' => 'rw'];
        $r_rw_document = ['group' => 'r', 'user' => 'r', 'search' => 'r', 'document' => 'rw'];
        $r_document = ['document' => 'r'];

        $this->call('config:cache', [], $buffer);
        print $buffer;

        $this->createDb();

        print "Running migrate\n";
        try {
            $this->call('migrate', ['--force' => null], $buffer);
            print $buffer;
        } catch (Exception $ex) {
            print $ex->getMessage() . "\n";
        }

        $url = config('app.urls.subscribe');
        if ($url) {
            printf("Subscribe to %s\n", $url);
            $response = Http::get($url);
            print_r($response->json());
        }

        $data_groups = [
            [
                'name' => 'Administradores',
                'privileges' => $rw,
                'authorities' => Group::makeAuthorities($rw),
                'types' => [0],
                'departments' => [0],
                'custom' => [],
                'searches' => [0],
            ],
            [
                'name' => 'Usuários',
                'privileges' => $r_rw_document,
                'authorities' => Group::makeAuthorities($r_rw_document),
                'types' => [0],
                'departments' => [0],
                'custom' => [],
                'searches' => [0],
            ],
            [
                'name' => 'Clientes',
                'privileges' => $r_document,
                'authorities' => Group::makeAuthorities($r_document),
                'types' => [0],
                'departments' => [0],
                'custom' => [['documents' => 'identity', 'users' => 'identity']],
                'searches' => [0],
            ]
        ];
        $groups = $this->createGroups($data_groups);
        $departments = $this->createDepartments([['name' => 'System'], ['name' => 'Administradores'], ['name' => 'Clientes']]);
        $this->createDocumentTypes([['name' => 'Documentos Públicos']]);


        $in_users = [
            [
                'data' => [
                    'name' => 'Sistema',
                    'department_id' => $departments['System'],
                    'identity' => '',
                    'phone' => '',
                    'email' => 'postmaster@localhost',
                    'username' => 'system',
                    'password' => config('app.passwords.system')
                ],
                'group' => [$groups['Administradores']]
            ],
            [
                'data' => [
                    'name' => 'Administrador',
                    'identity' => '',
                    'department_id' => $departments['System'],
                    'phone' => '',
                    'email' => 'admin@localhost',
                    'username' => 'admin',
                    'password' => config('app.passwords.admin')
                ],
                'group' => [$groups['Administradores']]
            ],
            [
                'data' => [
                    'name' => config('app.path'),
                    'identity' => '',
                    'department_id' => $departments['Administradores'],
                    'phone' => '',
                    'email' => 'client@localhost',
                    'username' => config('app.path'),
                    'password' => config('app.passwords.user')
                ],
                'group' => [$groups['Administradores']]
            ],
        ];
        $this->createUsers($in_users);

        $in_searches = [
            [
                'name' => 'Universal',
                'show_field' => [
                    'document_type' => true,
                    'department' => true,
                    'code' => true,
                    'identity' => true,
                    'name' => true,
                    'comment' => true,
                    'storage' => true,
                ]
            ]
        ];
        $this->createSearches($in_searches);
    }

    private function createDb()
    {

        $connection = config("database.default", "pgsql");
        $section = "database.connections." . $connection;
        $database = config($section . ".database");
        $username = config($section . ".username");
        $password = config($section . ".password");

        $querys = [];

        if ($connection == 'pgsql') {
            $querys['create_user'] = "CREATE USER {$username} WITH PASSWORD '$password'";
            $querys['create_db'] = "CREATE DATABASE $database OWNER $username";
            $querys['grant'] = "GRANT ALL PRIVILEGES ON DATABASE $database TO $username";
        } else {
            $querys['create_user'] = "CREATE USER '$username'@'%' IDENTIFIED BY '$password'";
            $querys['create_db'] = "CREATE DATABASE $database";
            $querys['grant'] = "GRANT PRIVILEGE ON $database TO '$database'@'%s'";
        }

        foreach ($querys as $q) {
            $show = preg_replace("/('$password')$/", "'******'", $q);
            printf("RUN [%s] ... ", $show);
            try {
                DB::connection($connection . ".super")->statement($q);
                print("[OK]\n");
            } catch (Exception $ex) {
                printf("[ERROR]\n[%s]\n", $ex->getMessage());
            }
        }
    }

    private function createGroups(array $in)
    {
        $out = [];
        foreach ($in as $item) {
            printf("Creating group [%s] ... ", $item['name']);
            $data = Group::where('name', $item['name'])->first();
            if ($data) {
                printf("[exists]\n", $item['name']);
                $out[$item['name']] = $data->id;
                continue;
            }
            try {
                $data = Group::create($item);
                $out[$item['name']] = $data->id;
                print("[created]\n");
            } catch (Exception $ex) {
                printf("[error]\n---\n%s\n---\n", $ex->getMessage());
            }
        }
        return $out;
    }



    private function createDocumentTypes(array $in)
    {
        $out = [];
        foreach ($in as $item) {
            printf("Creating document type [%s] ... ", $item['name']);
            $data = DocumentType::where('name', $item['name'])->first();
            if ($data) {
                printf("[exists]\n", $item['name']);
                $out[$item['name']] = $data->id;
                continue;
            }
            try {
                $data = DocumentType::create($item);
                $out[$item['name']] = $data->id;
                print("[created]\n");
            } catch (Exception $ex) {
                printf("[error]\n---\n%s\n---\n", $ex->getMessage());
            }
        }
        return $out;
    }

    private function createDepartments(array $in)
    {
        $out = [];
        foreach ($in as $item) {
            printf("Creating department [%s] ... ", $item['name']);
            $data = Department::where('name', $item['name'])->first();
            if ($data) {
                printf("[exists]\n");
                $out[$item['name']] = $data->id;
                continue;
            }
            try {
                $data = Department::create($item);
                $out[$item['name']] = $data->id;
                print("[created]\n");
            } catch (Exception $ex) {
                printf("[error]\n---\n%s\n---\n", $ex->getMessage());
            }
        }
        return $out;
    }


    private function createUsers(array $in)
    {
        $out = [];
        foreach ($in as $item) {
            $subitem = $item['data'];
            printf("Creating user [%s] ... ", $subitem['username']);
            $data = User::where('username', $subitem['username'])->first();
            if ($data) {
                printf("[exists]\n");
                $out[$subitem['username']] = $data->id;
                continue;
            }
            try {
                $data = User::create($subitem);
                $data->groups()->attach($item['group']);
                $out[$subitem['username']] = $data->id;
                print("[created]\n");
            } catch (Exception $ex) {
                printf("[error]\n---\n%s\n---\n", $ex->getMessage());
            }
        }
        return $out;
    }

    private function createSearches(array $in)
    {
        $out = [];
        foreach ($in as $item) {
            printf("Creating search [%s] ... ", $item['name']);
            $data = Search::where('name', $item['name'])->first();
            if ($data) {
                printf("[exists]\n");
                $out[$item['name']] = $data->id;
                continue;
            }
            try {
                $data = Search::create($item);
                $out[$item['name']] = $data->id;
                print("[created]\n");
            } catch (Exception $ex) {
                printf("[error]\n---\n%s\n---\n", $ex->getMessage());
            }
        }
        return $out;
    }
}
